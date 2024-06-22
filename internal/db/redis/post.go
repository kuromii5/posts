package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kuromii5/posts/internal/models"
	"github.com/redis/go-redis/v9"
)

func (d *RedisDB) SavePost(ctx context.Context, post *models.Post) error {
	const f = "redis.SavePost"

	// generate unique ID
	postID, err := d.client.Incr(ctx, "post:id").Result()
	if err != nil {
		return fmt.Errorf("%s: failed to increment post ID: %w", f, err)
	}
	post.ID = uint64(postID)

	postData, err := json.Marshal(post)
	if err != nil {
		return fmt.Errorf("%s: failed to marshal post: %w", f, err)
	}

	key := fmt.Sprintf("post:%d", post.ID)
	err = d.client.Set(ctx, key, postData, 0).Err()
	if err != nil {
		return fmt.Errorf("%s: failed to save post: %w", f, err)
	}

	return nil
}

func (d *RedisDB) PostByID(ctx context.Context, id uint64) (*models.Post, error) {
	const f = "redis.PostByID"

	key := fmt.Sprintf("post:%d", id)

	data, err := d.client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("%s: post not found", f)
		}
		return nil, fmt.Errorf("%s: failed to get post: %w", f, err)
	}

	var post models.Post
	err = json.Unmarshal(data, &post)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to unmarshal post data: %w", f, err)
	}

	return &post, nil
}

func (d *RedisDB) Posts(ctx context.Context) ([]*models.Post, error) {
	const f = "redis.Posts"

	var posts []*models.Post

	keys, err := d.client.Keys(ctx, "post:*").Result()
	if err != nil {
		return nil, fmt.Errorf("%s: failed to retrieve keys: %w", f, err)
	}

	for _, key := range keys {
		if key == "post:id" {
			continue
		}

		data, err := d.client.Get(ctx, key).Bytes()
		if err != nil {
			return nil, fmt.Errorf("%s: failed to get post data for key %s: %w", f, key, err)
		}

		var post models.Post
		err = json.Unmarshal(data, &post)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to unmarshal post data for key %s: %w", f, key, err)
		}

		posts = append(posts, &post)
	}

	return posts, nil
}
