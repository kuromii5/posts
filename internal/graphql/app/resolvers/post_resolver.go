package resolvers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/kuromii5/posts/internal/models"
)

type postResolver struct{ *Resolver }

// ID is the resolver for the id field.
func (r *postResolver) ID(ctx context.Context, obj *models.Post) (string, error) {
	return strconv.FormatUint(obj.ID, 10), nil
}

// User is the resolver for the user field.
func (r *postResolver) User(ctx context.Context, obj *models.Post) (*models.User, error) {
	if obj.User == nil {
		user, err := r.Service.UserByID(ctx, obj.UserID)
		if err != nil {
			return nil, fmt.Errorf("error fetching user: %v", err)
		}

		obj.User = user
	}
	return obj.User, nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *postResolver) CreatedAt(ctx context.Context, obj *models.Post) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *postResolver) UpdatedAt(ctx context.Context, obj *models.Post) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
}

// Comments is the resolver for the comments field.
func (r *postResolver) Comments(ctx context.Context, obj *models.Post, limit *int, offset *int) ([]*models.Comment, error) {
	if obj.Comments == nil {
		comments, err := r.Service.CommentsByPostID(ctx, obj.ID, limit, offset)
		if err != nil {
			return nil, fmt.Errorf("error fetching comments: %v", err)
		}

		obj.Comments = comments
	}
	return obj.Comments, nil
}
