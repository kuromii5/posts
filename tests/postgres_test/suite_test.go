package postgres_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	pg "github.com/kuromii5/posts/internal/db/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDBTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	db   *gorm.DB
	pgDB *pg.PostgresDB
}

func (suite *PostgresDBTestSuite) SetupSuite() {
	var err error

	// Create a mock database connection
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)

	// Initialize Gorm with the mock database connection
	suite.db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(suite.T(), err)

	suite.mock = mock

	suite.pgDB = &pg.PostgresDB{DB: suite.db}
}

func (suite *PostgresDBTestSuite) TearDownSuite() {
	sqlDB, err := suite.db.DB()
	assert.NoError(suite.T(), err)
	sqlDB.Close()
}

func (suite *PostgresDBTestSuite) TearDownTest() {
	suite.db.Rollback()
}

func TestPostgresDBTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresDBTestSuite))
}
