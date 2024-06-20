package resolvers

import (
	"context"
	"fmt"

	"github.com/kuromii5/posts/internal/models"
)

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	panic(fmt.Errorf("not implemented: Posts - posts"))
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*models.Post, error) {
	panic(fmt.Errorf("not implemented: Post - post"))
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, postID string) ([]*models.Comment, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}
