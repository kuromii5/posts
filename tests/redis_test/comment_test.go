package redis_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/kuromii5/posts/internal/db"
	"github.com/kuromii5/posts/internal/models"
	"github.com/kuromii5/posts/tests/utils"
	"github.com/stretchr/testify/assert"
)

func (suite *RedisDBTestSuite) TestSaveComment() {
	assert := assert.New(suite.T())

	comment := &models.Comment{
		Content:   "Test Comment",
		PostID:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := suite.redisDB.SaveComment(context.Background(), comment)
	assert.NoError(err)

	key := fmt.Sprintf("comment:%d", comment.ID)
	data, err := suite.redisDB.Client.Get(context.Background(), key).Result()
	assert.NoError(err)

	var savedComment models.Comment
	err = json.Unmarshal([]byte(data), &savedComment)
	assert.NoError(err)
	assert.Equal(comment.Content, savedComment.Content)
	assert.Equal(comment.PostID, savedComment.PostID)
	assert.Equal(comment.ID, savedComment.ID)
	assert.False(savedComment.CreatedAt.IsZero())
	assert.False(savedComment.UpdatedAt.IsZero())
}

func (suite *RedisDBTestSuite) TestCommentByID() {
	assert := assert.New(suite.T())

	comment := &models.Comment{
		Content:   "Test Comment",
		PostID:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := suite.redisDB.SaveComment(context.Background(), comment)
	assert.NoError(err)

	fetchedComment, err := suite.redisDB.CommentByID(context.Background(), comment.ID)
	assert.NoError(err)
	assert.Equal(comment.Content, fetchedComment.Content)
	assert.Equal(comment.PostID, fetchedComment.PostID)
	assert.Equal(comment.ID, fetchedComment.ID)
	assert.False(fetchedComment.CreatedAt.IsZero())
	assert.False(fetchedComment.UpdatedAt.IsZero())

	// Test for non-existent comment
	_, err = suite.redisDB.CommentByID(context.Background(), 999)
	assert.Error(err)
	assert.True(errors.Is(err, db.ErrNotFound))
}

func (suite *RedisDBTestSuite) TestCommentsByPostID() {
	assert := assert.New(suite.T())

	// Clear any existing comments
	err := suite.redisDB.Client.FlushDB(context.Background()).Err()
	assert.NoError(err)

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
		assert.NoError(err)
	}

	// Retrieve comments for PostID 1
	retrievedComments, err := suite.redisDB.CommentsByPostID(context.Background(), 1, 10, 0)
	assert.NoError(err)
	assert.Equal(2, len(retrievedComments))

	for _, comment := range retrievedComments {
		assert.Equal(uint64(1), comment.PostID)
		assert.False(comment.CreatedAt.IsZero())
		assert.False(comment.UpdatedAt.IsZero())
	}
}

func (suite *RedisDBTestSuite) TestRepliesByCommentID() {
	assert := assert.New(suite.T())

	// Clear any existing comments
	err := suite.redisDB.Client.FlushDB(context.Background()).Err()
	assert.NoError(err)

	parentComment := utils.CreateTestComment(1, 1, "Parent comment")

	// Save parent comment
	err = suite.redisDB.SaveComment(context.Background(), parentComment)
	assert.NoError(err)

	// Create and save multiple replies
	replies := []*models.Comment{
		utils.CreateTestReply(2, 1, 1, "Reply 1"),
		utils.CreateTestReply(2, 1, 1, "Reply 2"),
		utils.CreateTestComment(3, 1, "Non-reply Comment"),
	}
	for _, reply := range replies {
		err := suite.redisDB.SaveComment(context.Background(), reply)
		assert.NoError(err)
	}

	// Retrieve replies for the parent comment
	retrievedReplies, err := suite.redisDB.RepliesByCommentID(context.Background(), parentComment.ID, 10, 0)
	assert.NoError(err)
	assert.Equal(2, len(retrievedReplies))

	for _, reply := range retrievedReplies {
		assert.Equal(parentComment.ID, *reply.ParentCommentID)
		assert.False(reply.CreatedAt.IsZero())
		assert.False(reply.UpdatedAt.IsZero())
	}
}
