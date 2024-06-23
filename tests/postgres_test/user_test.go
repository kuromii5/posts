package postgres_test

import (
	"context"
	"errors"

	"github.com/kuromii5/posts/internal/db"
	"github.com/kuromii5/posts/internal/models"
	"github.com/kuromii5/posts/tests/utils"
	"github.com/stretchr/testify/assert"
)

func (suite *PostgresDBTestSuite) TestSaveUser() {
	assert := assert.New(suite.T())

	user := utils.CreateTestUser("testuser")
	err := suite.db.SaveUser(context.Background(), user)
	assert.NoError(err)
	assert.NotZero(user.ID)

	var savedUser models.User
	err = suite.db.DB.First(&savedUser, user.ID).Error
	assert.NoError(err)
	assert.Equal(user.Username, savedUser.Username)
}

func (suite *PostgresDBTestSuite) TestUserByID() {
	assert := assert.New(suite.T())

	user := utils.CreateTestUser("testuser")
	_ = suite.db.SaveUser(context.Background(), user)

	retrievedUser, err := suite.db.UserByID(context.Background(), user.ID)
	assert.NoError(err)
	assert.NotNil(retrievedUser)
	assert.Equal(user.Username, retrievedUser.Username)

	// Test for non-existent user
	_, err = suite.db.UserByID(context.Background(), 999)
	assert.Error(err)
	assert.True(errors.Is(err, db.ErrNotFound))
}
