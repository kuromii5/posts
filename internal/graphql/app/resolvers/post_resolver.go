package resolvers

import (
	"context"
	"fmt"

	"github.com/kuromii5/posts/internal/models"
)

// ID is the resolver for the id field.
func (r *postResolver) ID(ctx context.Context, obj *models.Post) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// User is the resolver for the user field.
func (r *postResolver) User(ctx context.Context, obj *models.Post) (*models.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// CreatedAt is the resolver for the createdAt field.
func (r *postResolver) CreatedAt(ctx context.Context, obj *models.Post) (string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - createdAt"))
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *postResolver) UpdatedAt(ctx context.Context, obj *models.Post) (string, error) {
	panic(fmt.Errorf("not implemented: UpdatedAt - updatedAt"))
}

// Comments is the resolver for the comments field.
func (r *postResolver) Comments(ctx context.Context, obj *models.Post) ([]*models.Comment, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}
