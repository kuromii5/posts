package redis_test

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kuromii5/posts/internal/models"
	"github.com/stretchr/testify/assert"
)

func (suite *RedisDBTestSuite) TestSavePost() {
	post := &models.Post{
		Title:     "Test Post",
		Content:   "This is a test post",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := suite.redisDB.SavePost(context.Background(), post)
	assert.NoError(suite.T(), err)

	key := fmt.Sprintf("post:%d", post.ID)
	data, err := suite.redisDB.Client.Get(context.Background(), key).Result()
	assert.NoError(suite.T(), err)

	var savedPost models.Post
	err = json.Unmarshal([]byte(data), &savedPost)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), post.Title, savedPost.Title)
	assert.Equal(suite.T(), post.Content, savedPost.Content)
	assert.Equal(suite.T(), post.ID, savedPost.ID)
	assert.False(suite.T(), savedPost.CreatedAt.IsZero())
	assert.False(suite.T(), savedPost.UpdatedAt.IsZero())
}

func (suite *RedisDBTestSuite) TestPostByID() {
	post := &models.Post{
		Title:     "Test Post",
		Content:   "This is a test post",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := suite.redisDB.SavePost(context.Background(), post)
	assert.NoError(suite.T(), err)

	fetchedPost, err := suite.redisDB.PostByID(context.Background(), post.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), post.Title, fetchedPost.Title)
	assert.Equal(suite.T(), post.Content, fetchedPost.Content)
	assert.Equal(suite.T(), post.ID, fetchedPost.ID)
	assert.False(suite.T(), fetchedPost.CreatedAt.IsZero())
	assert.False(suite.T(), fetchedPost.UpdatedAt.IsZero())
}

func (suite *RedisDBTestSuite) TestPosts() {
	// Clear any existing posts
	err := suite.redisDB.Client.FlushDB(context.Background()).Err()
	assert.NoError(suite.T(), err)

	// Create and save multiple posts
	posts := []*models.Post{
		{
			Title:     "Post 1",
			Content:   "Content 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Title:     "Post 2",
			Content:   "Content 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Title:     "Post 3",
			Content:   "Content 3",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, post := range posts {
		err := suite.redisDB.SavePost(context.Background(), post)
		assert.NoError(suite.T(), err)
	}

	// Retrieve all posts
	retrievedPosts, err := suite.redisDB.Posts(context.Background())
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), len(posts), len(retrievedPosts))

	for i, post := range posts {
		assert.Equal(suite.T(), post.Title, retrievedPosts[i].Title)
		assert.Equal(suite.T(), post.Content, retrievedPosts[i].Content)
		assert.Equal(suite.T(), post.ID, retrievedPosts[i].ID)
		assert.False(suite.T(), retrievedPosts[i].CreatedAt.IsZero())
		assert.False(suite.T(), retrievedPosts[i].UpdatedAt.IsZero())
	}
}
