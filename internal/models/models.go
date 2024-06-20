package models

import (
	"time"
)

type User struct {
	ID           uint64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Username     string     `json:"username"`
	Email        string     `json:"email" gorm:"uniqueIndex"`
	PasswordHash string     `json:"-"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Posts        []*Post    `json:"posts" gorm:"foreignKey:UserID"`
	Comments     []*Comment `json:"comments" gorm:"foreignKey:UserID"`
}

type Post struct {
	ID              uint64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID          uint64     `json:"user_id" gorm:"index"`
	Title           string     `json:"title"`
	Content         string     `json:"content"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	CommentsEnabled bool       `json:"comments_enabled"`
	User            *User      `json:"user" gorm:"foreignKey:UserID"`
	Comments        []*Comment `json:"comments" gorm:"foreignKey:PostID"`
}

type Comment struct {
	ID              uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	PostID          uint64    `json:"post_id" gorm:"index"`
	ParentCommentID *uint64   `json:"parent_comment_id,omitempty"`
	UserID          uint64    `json:"user_id" gorm:"index"`
	Content         string    `json:"content"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	User            *User     `json:"user" gorm:"foreignKey:UserID"`
	Post            *Post     `json:"post" gorm:"foreignKey:PostID"`
	ParentComment   *Comment  `json:"parent_comment,omitempty" gorm:"foreignKey:ParentCommentID"`
}
