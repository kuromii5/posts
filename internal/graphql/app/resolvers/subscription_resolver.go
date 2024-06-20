package resolvers

import (
	"context"
	"fmt"

	"github.com/kuromii5/posts/internal/models"
)

// CommentAdded is the resolver for the commentAdded field.
func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID string) (<-chan *models.Comment, error) {
	panic(fmt.Errorf("not implemented: CommentAdded - commentAdded"))
}
