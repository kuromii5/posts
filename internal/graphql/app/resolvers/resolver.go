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

func (r *Resolver) Comment() graph.CommentResolver { return &commentResolver{r} }

func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Post() graph.PostResolver { return &postResolver{r} }

func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

func (r *Resolver) Subscription() graph.SubscriptionResolver { return &subscriptionResolver{r} }

func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }
