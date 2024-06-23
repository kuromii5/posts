package postgres_test

import (
	"context"
	"errors"

	"github.com/kuromii5/posts/internal/db"
	"github.com/kuromii5/posts/tests/utils"
	"github.com/stretchr/testify/assert"
)

func (suite *PostgresDBTestSuite) TestSavePost() {
	assert := assert.New(suite.T())

	user := utils.CreateTestUser("testuser")
	err := suite.db.SaveUser(context.Background(), user)
	assert.NoError(err)
	assert.NotZero(user.ID)

	post := utils.CreateTestPost(user.ID, "Test Title", "Test Content", false)
	err = suite.db.SavePost(context.Background(), post)
	assert.NoError(err)
	assert.NotZero(post.ID)
}

func (suite *PostgresDBTestSuite) TestPostByID() {
	assert := assert.New(suite.T())

	user := utils.CreateTestUser("testuser")
	_ = suite.db.SaveUser(context.Background(), user)

	post := utils.CreateTestPost(user.ID, "Test Title", "Test Content", false)
	_ = suite.db.SavePost(context.Background(), post)

	retrievedPost, err := suite.db.PostByID(context.Background(), post.ID)
	assert.NoError(err)
	assert.NotNil(retrievedPost)
	assert.Equal(post.Title, retrievedPost.Title)

	// Test for non-existent post
	_, err = suite.db.PostByID(context.Background(), 999)
	assert.Error(err)
	assert.True(errors.Is(err, db.ErrNotFound))
}
