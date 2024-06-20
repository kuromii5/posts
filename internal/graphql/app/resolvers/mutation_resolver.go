package resolvers

import (
	"context"
	"errors"
	"fmt"

	model "github.com/kuromii5/posts/internal/graphql/app/domain"
	"github.com/kuromii5/posts/internal/models"
	"github.com/kuromii5/posts/internal/service"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*models.User, error) {
	user, err := r.Service.CreateUser(ctx, input.Username, input.Email, input.Password)
	if err != nil {
		if errors.Is(err, service.ErrUserExists) {
			return nil, fmt.Errorf("user already exists")
		}
		return nil, fmt.Errorf("internal server error")
	}

	return user, nil
}

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUser) (*model.AuthPayload, error) {
	token, err := r.Service.LoginUser(ctx, input.Email, input.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, fmt.Errorf("invalid credentials")
		}
		return nil, fmt.Errorf("internal server error")
	}

	// Prepare and return auth payload
	authPayload := &model.AuthPayload{Token: token}
	return authPayload, nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*models.Post, error) {
	panic(fmt.Errorf("not implemented: CreatePost - createPost"))
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*models.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - createComment"))
}

// TriggerPostCommentState is the resolver for the triggerPostCommentState field.
func (r *mutationResolver) TriggerPostCommentState(ctx context.Context, input model.PostCommentState) (*models.Post, error) {
	panic(fmt.Errorf("not implemented: TriggerPostCommentState - triggerPostCommentState"))
}
