package postgres_test

import (
	"testing"

	pg "github.com/kuromii5/posts/internal/db/postgres"
	"github.com/kuromii5/posts/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDBTestSuite struct {
	suite.Suite
	db *pg.PostgresDB
}

var dbUrl = "postgres://postgres:admin@localhost:5432/reddit_clone_test?sslmode=disable"

func (suite *PostgresDBTestSuite) SetupSuite() {
	var err error

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	assert.NoError(suite.T(), err)
	suite.db = &pg.PostgresDB{DB: db}

	// Auto migrate the database schema for testing
	err = suite.db.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	assert.NoError(suite.T(), err)
}

func (suite *PostgresDBTestSuite) TearDownSuite() {
	sqlDB, err := suite.db.DB.DB()
	assert.NoError(suite.T(), err)
	sqlDB.Close()
}

func (suite *PostgresDBTestSuite) TearDownTest() {
	// Clean up the database after each test
	suite.db.DB.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
}

func TestPostgresDBTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresDBTestSuite))
}
