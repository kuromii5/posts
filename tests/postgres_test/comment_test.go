package postgres_test

import (
	"context"
	"errors"

	"github.com/kuromii5/posts/internal/db"
	"github.com/kuromii5/posts/tests/utils"
	"github.com/stretchr/testify/assert"
)

func (suite *PostgresDBTestSuite) TestSaveComment() {
	assert := assert.New(suite.T())

	user := utils.CreateTestUser("testuser")
	_ = suite.db.SaveUser(context.Background(), user)

	post := utils.CreateTestPost(user.ID, "Test Title", "Test Content", true)
	_ = suite.db.SavePost(context.Background(), post)

	comment := utils.CreateTestComment(user.ID, post.ID, "Test Comment")
	err := suite.db.SaveComment(context.Background(), comment)
	assert.NoError(err)
	assert.NotZero(comment.ID)
}

func (suite *PostgresDBTestSuite) TestCommentByID() {
	assert := assert.New(suite.T())

	user := utils.CreateTestUser("testuser")
	_ = suite.db.SaveUser(context.Background(), user)

	post := utils.CreateTestPost(user.ID, "Test Title", "Test Content", true)
	_ = suite.db.SavePost(context.Background(), post)

	comment := utils.CreateTestComment(user.ID, post.ID, "Test Comment")
	_ = suite.db.SaveComment(context.Background(), comment)

	retrievedComment, err := suite.db.CommentByID(context.Background(), comment.ID)
	assert.NoError(err)
	assert.NotNil(retrievedComment)
	assert.Equal(comment.Content, retrievedComment.Content)

	// Test for non-existent comment
	_, err = suite.db.CommentByID(context.Background(), 999)
	assert.Error(err)
	assert.True(errors.Is(err, db.ErrNotFound))
}

func (suite *PostgresDBTestSuite) TestCommentsByPostID() {
	assert := assert.New(suite.T())

	user := utils.CreateTestUser("testuser")
	_ = suite.db.SaveUser(context.Background(), user)

	post := utils.CreateTestPost(user.ID, "Test Title", "Test Content", true)
	_ = suite.db.SavePost(context.Background(), post)

	comment1 := utils.CreateTestComment(1, 1, "Test comment 1")
	comment2 := utils.CreateTestComment(1, 1, "Test comment 2")
	_ = suite.db.SaveComment(context.Background(), comment1)
	_ = suite.db.SaveComment(context.Background(), comment2)

	comments, err := suite.db.CommentsByPostID(context.Background(), 1, 10, 0)
	assert.NoError(err)
	assert.Len(comments, 2)
}

func (suite *PostgresDBTestSuite) TestRepliesByCommentID() {
	assert := assert.New(suite.T())

	user := utils.CreateTestUser("testuser")
	_ = suite.db.SaveUser(context.Background(), user)

	post := utils.CreateTestPost(user.ID, "Test Title", "Test Content", true)
	_ = suite.db.SavePost(context.Background(), post)

	comment := utils.CreateTestComment(1, 1, "Test comment")
	err := suite.db.SaveComment(context.Background(), comment)
	assert.NoError(err)
	assert.NotZero(comment.ID)

	reply1 := utils.CreateTestReply(1, 1, comment.ID, "Reply 1")
	reply2 := utils.CreateTestReply(1, 1, comment.ID, "Reply 2")

	err = suite.db.SaveComment(context.Background(), reply1)
	assert.NoError(err)
	err = suite.db.SaveComment(context.Background(), reply2)
	assert.NoError(err)

	replies, err := suite.db.RepliesByCommentID(context.Background(), comment.ID, 10, 0)
	assert.NoError(err)
	assert.Len(replies, 2)
}
