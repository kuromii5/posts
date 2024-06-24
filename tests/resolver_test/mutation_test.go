package resolver_test

import (
	"context"
	"testing"

	model "github.com/kuromii5/posts/internal/graphql/app/domain"
	"github.com/kuromii5/posts/internal/graphql/app/resolvers"
	"github.com/kuromii5/posts/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) CreateUser(ctx context.Context, username string) (*models.User, error) {
	args := m.Called(ctx, username)
	if user, ok := args.Get(0).(*models.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestCreateUser(t *testing.T) {
	// Создаем моки для сервисного слоя
	mockService := new(MockService)
	testResolver := resolvers.NewResolver(mockService)

	// Настраиваем ожидаемое поведение мока
	expectedUser := &models.User{
		Username: "testuser",
	}
	mockService.On("CreateUser", mock.Anything, "testuser").Return(expectedUser, nil)

	input := model.NewUser{
		Username: "testuser",
	}
	result, err := testResolver.CreateUser(context.Background(), input)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "testuser", result.Username)

	mockService.AssertCalled(t, "CreateUser", mock.Anything, "testuser")
}
