package redis_test

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kuromii5/posts/internal/models"
	"github.com/stretchr/testify/assert"
)

func (suite *RedisDBTestSuite) TestSaveAndGetUser() {
	user := &models.User{
		Username:  "testuser",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := suite.redisDB.SaveUser(context.Background(), user)
	assert.NoError(suite.T(), err)

	key := "user:" + fmt.Sprint(user.ID)
	data, err := suite.redisDB.Client.Get(context.Background(), key).Result()
	assert.NoError(suite.T(), err)

	var savedUser models.User
	err = json.Unmarshal([]byte(data), &savedUser)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), "testuser", savedUser.Username)
	assert.Equal(suite.T(), uint64(1), savedUser.ID)
	assert.False(suite.T(), savedUser.CreatedAt.IsZero())
	assert.False(suite.T(), savedUser.UpdatedAt.IsZero())
}

func (suite *RedisDBTestSuite) TestUserByID() {
	user := &models.User{
		Username:  "testuser",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := suite.redisDB.SaveUser(context.Background(), user)
	assert.NoError(suite.T(), err)

	fetchedUser, err := suite.redisDB.UserByID(context.Background(), user.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Username, fetchedUser.Username)
	assert.Equal(suite.T(), user.ID, fetchedUser.ID)
	assert.False(suite.T(), fetchedUser.CreatedAt.IsZero())
	assert.False(suite.T(), fetchedUser.UpdatedAt.IsZero())
}

func (suite *RedisDBTestSuite) TestUserByID_NotFound() {
	_, err := suite.redisDB.UserByID(context.Background(), 9999)
	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "not found")
}
