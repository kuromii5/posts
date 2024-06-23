package redis_test

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	r "github.com/kuromii5/posts/internal/db/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RedisDBTestSuite struct {
	suite.Suite
	redisServer *miniredis.Miniredis
	redisDB     *r.RedisDB
}

func (suite *RedisDBTestSuite) SetupSuite() {
	var err error

	suite.redisServer, err = miniredis.Run()
	assert.NoError(suite.T(), err)

	suite.redisDB, err = r.New(suite.redisServer.Addr())
	assert.NoError(suite.T(), err)
}

func (suite *RedisDBTestSuite) TearDownSuite() {
	suite.redisServer.Close()
}

func TestRedisDBTestSuite(t *testing.T) {
	suite.Run(t, new(RedisDBTestSuite))
}
