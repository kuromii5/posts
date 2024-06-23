package resolvers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kuromii5/posts/internal/models"
)

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID string) (<-chan *models.Comment, error) {
	postId, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid post ID")
	}

	commentsChan, unsubscribe := r.Service.PubSubService.Subscribe(postId)

	go func() {
		<-ctx.Done()
		unsubscribe()
	}()

	return commentsChan, nil
}
