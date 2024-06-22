package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kuromii5/posts/internal/models"
	"github.com/redis/go-redis/v9"
)

func (d *RedisDB) SaveComment(ctx context.Context, comment *models.Comment) error {
	const f = "redis.SaveComment"

	// generate unique ID
	commID, err := d.client.Incr(ctx, "comment:id").Result()
	if err != nil {
		return fmt.Errorf("%s: failed to increment comment ID: %w", f, err)
	}
	comment.ID = uint64(commID)

	commentData, err := json.Marshal(comment)
	if err != nil {
		return fmt.Errorf("%s: failed to marshal comment: %w", f, err)
	}

	key := fmt.Sprintf("comment:%d", comment.ID)

	err = d.client.Set(ctx, key, commentData, 0).Err()
	if err != nil {
		return fmt.Errorf("%s: failed to save comment: %w", f, err)
	}

	return nil
}

func (d *RedisDB) CommentByID(ctx context.Context, id uint64) (*models.Comment, error) {
	const f = "redis.CommentByID"

	key := fmt.Sprintf("comment:%d", id)

	data, err := d.client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("%s: comment not found", f)
		}
		return nil, fmt.Errorf("%s: failed to get comment: %w", f, err)
	}

	var comment models.Comment
	err = json.Unmarshal(data, &comment)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to unmarshal comment data: %w", f, err)
	}

	return &comment, nil
}

func (d *RedisDB) CommentsByPostID(ctx context.Context, postID uint64, limit, offset int) ([]*models.Comment, error) {
	const f = "redis.CommentsByPostID"

	var comments []*models.Comment

	keys, err := d.client.Keys(ctx, "comment:*").Result()
	if err != nil {
		return nil, fmt.Errorf("%s: failed to retrieve keys: %w", f, err)
	}

	for _, key := range keys {
		if key == "comment:id" {
			continue
		}

		data, err := d.client.Get(ctx, key).Bytes()
		if err != nil {
			return nil, fmt.Errorf("%s: failed to get comment data for key %s: %w", f, key, err)
		}

		var comment models.Comment
		err = json.Unmarshal(data, &comment)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to unmarshal comment data for key %s: %w", f, key, err)
		}

		// check if matches postID and it is not reply
		if comment.PostID == postID && comment.ParentCommentID == nil {
			comments = append(comments, &comment)
		}
	}

	// pagination
	start := offset
	end := offset + limit
	if start > len(comments) {
		start = len(comments)
	}
	if end > len(comments) {
		end = len(comments)
	}
	comments = comments[start:end]

	return comments, nil
}

func (d *RedisDB) RepliesByCommentID(ctx context.Context, commID uint64, limit, offset int) ([]*models.Comment, error) {
	const f = "redis.RepliesByCommentID"

	var replies []*models.Comment

	keys, err := d.client.Keys(ctx, "comment:*").Result()
	if err != nil {
		return nil, fmt.Errorf("%s: failed to retrieve keys: %w", f, err)
	}

	for _, key := range keys {
		if key == "comment:id" {
			continue
		}

		data, err := d.client.Get(ctx, key).Bytes()
		if err != nil {
			return nil, fmt.Errorf("%s: failed to get reply data for key %s: %w", f, key, err)
		}

		var reply models.Comment
		err = json.Unmarshal(data, &reply)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to unmarshal reply data for key %s: %w", f, key, err)
		}

		// Check if the current reply's ParentCommentID matches the commID
		if reply.ParentCommentID != nil && *reply.ParentCommentID == commID {
			replies = append(replies, &reply)
		}
	}

	// pagination
	start := offset
	end := offset + limit
	if start > len(replies) {
		start = len(replies)
	}
	if end > len(replies) {
		end = len(replies)
	}
	replies = replies[start:end]

	return replies, nil
}
