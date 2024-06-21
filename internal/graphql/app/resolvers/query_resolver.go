package resolvers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kuromii5/posts/internal/models"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %v", err)
	}

	user, err := r.Service.UserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return user, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, limit *int, offset *int) ([]*models.Post, error) {
	var posts []*models.Post
	var err error

	posts, err = r.Service.Posts(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %v", err)
	}

	return posts, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*models.Post, error) {
	postID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid post ID: %v", err)
	}

	post, err := r.Service.PostByID(ctx, postID)
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %v", err)
	}

	return post, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Comments(ctx context.Context, postID string, limit *int, offset *int) ([]*models.Comment, error) {
	postId, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid post ID: %v", err)
	}

	var comments []*models.Comment
	comments, err = r.Service.CommentsByPostID(ctx, postId, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments: %v", err)
	}

	return comments, nil
}
