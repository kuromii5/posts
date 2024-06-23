package redis_test

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kuromii5/posts/internal/db"
	"github.com/kuromii5/posts/internal/models"
	"github.com/kuromii5/posts/tests/utils"
	"github.com/stretchr/testify/assert"
)

func (suite *RedisDBTestSuite) TestSaveAndGetUser() {
	assert := assert.New(suite.T())

	user := utils.CreateTestUser("testuser")
	err := suite.redisDB.SaveUser(context.Background(), user)
	assert.NoError(err)
	assert.NotZero(user.ID)

	key := "user:" + fmt.Sprint(user.ID)
	data, err := suite.redisDB.Client.Get(context.Background(), key).Result()
	assert.NoError(err)

	var savedUser models.User
	err = json.Unmarshal([]byte(data), &savedUser)
	assert.NoError(err)

	assert.Equal("testuser", savedUser.Username)
	assert.Equal(uint64(1), savedUser.ID)
	assert.False(savedUser.CreatedAt.IsZero())
	assert.False(savedUser.UpdatedAt.IsZero())
}

func (suite *RedisDBTestSuite) TestUserByID() {
	assert := assert.New(suite.T())

	user := utils.CreateTestUser("testuser")
	err := suite.redisDB.SaveUser(context.Background(), user)
	assert.NoError(err)
	assert.NotZero(user.ID)

	fetchedUser, err := suite.redisDB.UserByID(context.Background(), user.ID)
	assert.NoError(err)
	assert.Equal(user.Username, fetchedUser.Username)
	assert.Equal(user.ID, fetchedUser.ID)
	assert.False(fetchedUser.CreatedAt.IsZero())
	assert.False(fetchedUser.UpdatedAt.IsZero())

	// Test for non-existent user
	_, err = suite.redisDB.UserByID(context.Background(), 9999)
	assert.Error(err)
	assert.Contains(err.Error(), db.ErrNotFound)
}
