package resolvers

import (
	"context"
	"fmt"
	"strconv"

	model "github.com/kuromii5/posts/internal/graphql/app/domain"
	"github.com/kuromii5/posts/internal/models"
)

type mutationResolver struct{ *Resolver }

// Save user in database
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	user, err := r.Service.CreateUser(ctx, input.Username)
	if err != nil {
		return nil, fmt.Errorf("error creating user")
	}

	return user, nil
}

// Save post in database
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*models.Post, error) {
	userID, err := strconv.ParseUint(input.UserID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID")
	}

	post, err := r.Service.CreatePost(ctx, input.Title, input.Content, userID, input.CommentsEnabled)
	if err != nil {
		return nil, fmt.Errorf("error creating post")
	}

	return post, nil
}

// Save comment in database
func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*models.Comment, error) {
	postID, err := strconv.ParseUint(input.PostID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid post ID")
	}

	userID, err := strconv.ParseUint(input.UserID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID")
	}

	// Check if there is parent comment
	// if so - bind id pointer to comment
	var parentCommentID *uint64
	if input.ParentCommentID != nil {
		id, err := strconv.ParseUint(*input.ParentCommentID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid parent comment ID")
		}
		parentCommentID = &id
	}

	// Fetch the post to check if comments are enabled
	post, err := r.Service.PostByID(ctx, postID)
	if err != nil {
		return nil, fmt.Errorf("error fetching post: %v", err)
	}

	// Check if comments are enabled for the post
	if !post.CommentsEnabled {
		return nil, fmt.Errorf("comments are not enabled for this post")
	}

	comment, err := r.Service.CreateComment(ctx, userID, postID, parentCommentID, input.Content)
	if err != nil {
		return nil, fmt.Errorf("error creating comment")
	}

	// if it is not reply - notify users about new comment
	if parentCommentID == nil {
		r.Service.PubSubService.Publish(postID, comment)
	}

	return comment, nil
}
