package resolvers

import (
	"context"
	"fmt"

	"github.com/kuromii5/posts/internal/models"
)

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// CreatedAt is the resolver for the createdAt field.
func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - createdAt"))
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *userResolver) UpdatedAt(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented: UpdatedAt - updatedAt"))
}
