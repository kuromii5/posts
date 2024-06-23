package utils

import (
	"time"

	"github.com/kuromii5/posts/internal/models"
)

func CreateTestUser(username string) *models.User {
	return &models.User{
		Username:  username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func CreateTestPost(userID uint64, title, content string, commentsEnabled bool) *models.Post {
	return &models.Post{
		UserID:          userID,
		Title:           title,
		Content:         content,
		CommentsEnabled: commentsEnabled,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func CreateTestComment(userID, postID uint64, content string) *models.Comment {
	return &models.Comment{
		UserID:    userID,
		PostID:    postID,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func CreateTestReply(userID, postID, parentCommentID uint64, content string) *models.Comment {
	return &models.Comment{
		UserID:          userID,
		PostID:          postID,
		ParentCommentID: &parentCommentID,
		Content:         content,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}
