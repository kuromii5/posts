package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kuromii5/posts/internal/db"
	"github.com/kuromii5/posts/internal/models"
	"github.com/redis/go-redis/v9"
)

func (d *RedisDB) SavePost(ctx context.Context, post *models.Post) error {
	const f = "redis.SavePost"

	// generate unique ID
	postID, err := d.Client.Incr(ctx, "post:id").Result()
	if err != nil {
		return fmt.Errorf("%s:%w", f, err)
	}
	post.ID = uint64(postID)

	postData, err := json.Marshal(post)
	if err != nil {
		return fmt.Errorf("%s:%w", f, err)
	}

	key := fmt.Sprintf("post:%d", post.ID)
	err = d.Client.Set(ctx, key, postData, 0).Err()
	if err != nil {
		return fmt.Errorf("%s:%w", f, err)
	}

	return nil
}

func (d *RedisDB) PostByID(ctx context.Context, id uint64) (*models.Post, error) {
	const f = "redis.PostByID"

	key := fmt.Sprintf("post:%d", id)

	data, err := d.Client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("%s:%w", f, db.ErrNotFound)
		}
		return nil, fmt.Errorf("%s:%w", f, err)
	}

	var post models.Post
	err = json.Unmarshal(data, &post)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", f, err)
	}

	return &post, nil
}

func (d *RedisDB) Posts(ctx context.Context) ([]*models.Post, error) {
	const f = "redis.Posts"

	var posts []*models.Post

	keys, err := d.Client.Keys(ctx, "post:*").Result()
	if err != nil {
		return nil, fmt.Errorf("%s:%w", f, err)
	}

	for _, key := range keys {
		if key == "post:id" {
			continue
		}

		data, err := d.Client.Get(ctx, key).Bytes()
		if err != nil {
			return nil, fmt.Errorf("%s:%w", f, err)
		}

		var post models.Post
		err = json.Unmarshal(data, &post)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", f, err)
		}

		posts = append(posts, &post)
	}

	return posts, nil
}
