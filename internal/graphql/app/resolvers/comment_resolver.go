package resolvers

import (
	"context"
	"fmt"

	"github.com/kuromii5/posts/internal/models"
)

// ID is the resolver for the id field.
func (r *commentResolver) ID(ctx context.Context, obj *models.Comment) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Post is the resolver for the post field.
func (r *commentResolver) Post(ctx context.Context, obj *models.Comment) (*models.Post, error) {
	panic(fmt.Errorf("not implemented: Post - post"))
}

// ParentComment is the resolver for the parentComment field.
func (r *commentResolver) ParentComment(ctx context.Context, obj *models.Comment) (*models.Comment, error) {
	panic(fmt.Errorf("not implemented: ParentComment - parentComment"))
}

// User is the resolver for the user field.
func (r *commentResolver) User(ctx context.Context, obj *models.Comment) (*models.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// CreatedAt is the resolver for the createdAt field.
func (r *commentResolver) CreatedAt(ctx context.Context, obj *models.Comment) (string, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - createdAt"))
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *commentResolver) UpdatedAt(ctx context.Context, obj *models.Comment) (string, error) {
	panic(fmt.Errorf("not implemented: UpdatedAt - updatedAt"))
}
