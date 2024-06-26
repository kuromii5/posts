// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewComment struct {
	PostID          string  `json:"postId"`
	ParentCommentID *string `json:"parentCommentId,omitempty"`
	UserID          string  `json:"userId"`
	Content         string  `json:"content"`
}

type NewPost struct {
	UserID          string `json:"userId"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	CommentsEnabled bool   `json:"commentsEnabled"`
}

type NewUser struct {
	Username string `json:"username"`
}

type Query struct {
}

type Subscription struct {
}
