package redis_test

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kuromii5/posts/internal/db"
	"github.com/kuromii5/posts/internal/models"
	"github.com/kuromii5/posts/tests/utils"
	"github.com/stretchr/testify/assert"
)

func (suite *RedisDBTestSuite) TestSavePost() {
	assert := assert.New(suite.T())

	post := &models.Post{
		Title:     "Test Post",
		Content:   "This is a test post",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := suite.redisDB.SavePost(context.Background(), post)
	assert.NoError(err)

	key := fmt.Sprintf("post:%d", post.ID)
	data, err := suite.redisDB.Client.Get(context.Background(), key).Result()
	assert.NoError(err)

	var savedPost models.Post
	err = json.Unmarshal([]byte(data), &savedPost)
	assert.NoError(err)
	assert.Equal(post.Title, savedPost.Title)
	assert.Equal(post.Content, savedPost.Content)
	assert.Equal(post.ID, savedPost.ID)
	assert.False(savedPost.CreatedAt.IsZero())
	assert.False(savedPost.UpdatedAt.IsZero())
}

func (suite *RedisDBTestSuite) TestPostByID() {
	assert := assert.New(suite.T())

	post := utils.CreateTestPost(1, "Test Post", "This is a test post", false)
	err := suite.redisDB.SavePost(context.Background(), post)
	assert.NoError(err)

	fetchedPost, err := suite.redisDB.PostByID(context.Background(), post.ID)
	assert.NoError(err)
	assert.Equal(post.Title, fetchedPost.Title)
	assert.Equal(post.Content, fetchedPost.Content)
	assert.Equal(post.ID, fetchedPost.ID)
	assert.False(fetchedPost.CreatedAt.IsZero())
	assert.False(fetchedPost.UpdatedAt.IsZero())

	// Test for non-existent post
	_, err = suite.redisDB.PostByID(context.Background(), 9999)
	assert.Error(err)
	assert.Contains(err.Error(), db.ErrNotFound)
}

func (suite *RedisDBTestSuite) TestPosts() {
	assert := assert.New(suite.T())

	// Clear any existing posts
	err := suite.redisDB.Client.FlushDB(context.Background()).Err()
	assert.NoError(err)

	// Create and save multiple posts
	posts := []*models.Post{
		utils.CreateTestPost(1, "Post 1", "Content 1", false),
		utils.CreateTestPost(1, "Post 2", "Content 2", false),
		utils.CreateTestPost(1, "Post 3", "Content 3", false),
	}
	for _, post := range posts {
		err := suite.redisDB.SavePost(context.Background(), post)
		assert.NoError(err)
	}

	// Retrieve all posts
	retrievedPosts, err := suite.redisDB.Posts(context.Background())
	assert.NoError(err)
	assert.Equal(len(posts), len(retrievedPosts))

	for i, post := range posts {
		assert.Equal(post.Title, retrievedPosts[i].Title)
		assert.Equal(post.Content, retrievedPosts[i].Content)
		assert.Equal(post.ID, retrievedPosts[i].ID)
		assert.False(retrievedPosts[i].CreatedAt.IsZero())
		assert.False(retrievedPosts[i].UpdatedAt.IsZero())
	}
}
