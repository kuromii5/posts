package redis_test

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kuromii5/posts/internal/models"
	"github.com/stretchr/testify/assert"
)

func (suite *RedisDBTestSuite) TestSaveComment() {
	comment := &models.Comment{
		Content:   "Test Comment",
		PostID:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := suite.redisDB.SaveComment(context.Background(), comment)
	assert.NoError(suite.T(), err)

	key := fmt.Sprintf("comment:%d", comment.ID)
	data, err := suite.redisDB.Client.Get(context.Background(), key).Result()
	assert.NoError(suite.T(), err)

	var savedComment models.Comment
	err = json.Unmarshal([]byte(data), &savedComment)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), comment.Content, savedComment.Content)
	assert.Equal(suite.T(), comment.PostID, savedComment.PostID)
	assert.Equal(suite.T(), comment.ID, savedComment.ID)
	assert.False(suite.T(), savedComment.CreatedAt.IsZero())
	assert.False(suite.T(), savedComment.UpdatedAt.IsZero())
}

func (suite *RedisDBTestSuite) TestCommentByID() {
	comment := &models.Comment{
		Content:   "Test Comment",
		PostID:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := suite.redisDB.SaveComment(context.Background(), comment)
	assert.NoError(suite.T(), err)

	fetchedComment, err := suite.redisDB.CommentByID(context.Background(), comment.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), comment.Content, fetchedComment.Content)
	assert.Equal(suite.T(), comment.PostID, fetchedComment.PostID)
	assert.Equal(suite.T(), comment.ID, fetchedComment.ID)
	assert.False(suite.T(), fetchedComment.CreatedAt.IsZero())
	assert.False(suite.T(), fetchedComment.UpdatedAt.IsZero())
}

func (suite *RedisDBTestSuite) TestCommentsByPostID() {
	// Clear any existing comments
	err := suite.redisDB.Client.FlushDB(context.Background()).Err()
	assert.NoError(suite.T(), err)

	// Create and save multiple comments
	comments := []*models.Comment{
		{
			Content:   "Comment 1",
			PostID:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Content:   "Comment 2",
			PostID:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Content:   "Comment 3",
			PostID:    2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, comment := range comments {
		err := suite.redisDB.SaveComment(context.Background(), comment)
		assert.NoError(suite.T(), err)
	}

	// Retrieve comments for PostID 1
	retrievedComments, err := suite.redisDB.CommentsByPostID(context.Background(), 1, 10, 0)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 2, len(retrievedComments))

	for _, comment := range retrievedComments {
		assert.Equal(suite.T(), uint64(1), comment.PostID)
		assert.False(suite.T(), comment.CreatedAt.IsZero())
		assert.False(suite.T(), comment.UpdatedAt.IsZero())
	}
}

func (suite *RedisDBTestSuite) TestRepliesByCommentID() {
	// Clear any existing comments
	err := suite.redisDB.Client.FlushDB(context.Background()).Err()
	assert.NoError(suite.T(), err)

	parentComment := &models.Comment{
		Content:   "Parent Comment",
		PostID:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save parent comment
	err = suite.redisDB.SaveComment(context.Background(), parentComment)
	assert.NoError(suite.T(), err)

	// Create and save multiple replies
	replies := []*models.Comment{
		{
			Content:         "Reply 1",
			PostID:          1,
			ParentCommentID: &parentComment.ID,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			Content:         "Reply 2",
			PostID:          1,
			ParentCommentID: &parentComment.ID,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			Content:   "Non-reply Comment",
			PostID:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, reply := range replies {
		err := suite.redisDB.SaveComment(context.Background(), reply)
		assert.NoError(suite.T(), err)
	}

	// Retrieve replies for the parent comment
	retrievedReplies, err := suite.redisDB.RepliesByCommentID(context.Background(), parentComment.ID, 10, 0)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 2, len(retrievedReplies))

	for _, reply := range retrievedReplies {
		assert.Equal(suite.T(), parentComment.ID, *reply.ParentCommentID)
		assert.False(suite.T(), reply.CreatedAt.IsZero())
		assert.False(suite.T(), reply.UpdatedAt.IsZero())
	}
}
