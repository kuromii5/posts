package resolvers

import (
	"github.com/kuromii5/posts/internal/graphql/graph"
	"github.com/kuromii5/posts/internal/service"
)

// Resolver struct will hold your dependencies, like database connections
//
//go:generate go run github.com/99designs/gqlgen
type Resolver struct {
	Service service.Service
}

// Comment returns graph.CommentResolver implementation.
func (r *Resolver) Comment() graph.CommentResolver { return &commentResolver{r} }

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Post returns graph.PostResolver implementation.
func (r *Resolver) Post() graph.PostResolver { return &postResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// Subscription returns graph.SubscriptionResolver implementation.
func (r *Resolver) Subscription() graph.SubscriptionResolver { return &subscriptionResolver{r} }

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

type commentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
