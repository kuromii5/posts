package resolvers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/kuromii5/posts/internal/models"
)

type commentResolver struct{ *Resolver }

// ID is the resolver for the id field.
func (r *commentResolver) ID(ctx context.Context, obj *models.Comment) (string, error) {
	return strconv.FormatUint(obj.ID, 10), nil
}

// Post is the resolver for the post field.
func (r *commentResolver) Post(ctx context.Context, obj *models.Comment) (*models.Post, error) {
	if obj.Post == nil {
		post, err := r.Service.PostByID(ctx, obj.PostID)
		if err != nil {
			return nil, fmt.Errorf("error fetching post: %v", err)
		}

		obj.Post = post
	}
	return obj.Post, nil
}

// ParentComment is the resolver for the parentComment field.
func (r *commentResolver) ParentComment(ctx context.Context, obj *models.Comment) (*models.Comment, error) {
	if obj.ParentCommentID == nil {
		return nil, nil
	}

	parentComment, err := r.Service.CommentByID(ctx, *obj.ParentCommentID)
	if err != nil {
		return nil, fmt.Errorf("error fetching parent comment: %v", err)
	}

	obj.ParentComment = parentComment
	return obj.ParentComment, nil
}

func (r *commentResolver) User(ctx context.Context, obj *models.Comment) (*models.User, error) {
	if obj.User == nil {
		user, err := r.Service.UserByID(ctx, obj.UserID)
		if err != nil {
			return nil, fmt.Errorf("error fetching user: %v", err)
		}

		obj.User = user
	}

	return obj.User, nil
}

// Replies is the resolver for the replies field.
func (r *commentResolver) Replies(ctx context.Context, obj *models.Comment, limit, offset *int) ([]*models.Comment, error) {
	replies, err := r.Service.RepliesByCommentID(ctx, obj.ID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error fetching replies: %v", err)
	}
	return replies, nil
}

func (r *commentResolver) CreatedAt(ctx context.Context, obj *models.Comment) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *commentResolver) UpdatedAt(ctx context.Context, obj *models.Comment) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
}
