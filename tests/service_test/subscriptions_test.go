package service_test

import (
	"testing"
	"time"

	"github.com/kuromii5/posts/internal/lib/logger"
	"github.com/kuromii5/posts/internal/models"
	"github.com/kuromii5/posts/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptionManager(t *testing.T) {
	manager := service.NewSubscriptionManager(logger.New(""))

	postID := uint64(1)
	comment := &models.Comment{
		ID:        1,
		PostID:    postID,
		UserID:    1,
		Content:   "Test comment",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	ch, unsubscribe := manager.Subscribe(postID)
	defer unsubscribe()

	go manager.Publish(postID, comment)

	receivedComment := <-ch

	assert.NotNil(t, receivedComment)
	assert.Equal(t, comment.ID, receivedComment.ID)
	assert.Equal(t, comment.PostID, receivedComment.PostID)
	assert.Equal(t, comment.UserID, receivedComment.UserID)
	assert.Equal(t, comment.Content, receivedComment.Content)
	assert.Equal(t, comment.CreatedAt.Unix(), receivedComment.CreatedAt.Unix())
	assert.Equal(t, comment.UpdatedAt.Unix(), receivedComment.UpdatedAt.Unix())
}
