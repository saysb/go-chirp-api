// internal/models/comment.go
package models

import (
	"errors"
	"time"
)

type Comment struct {
    ID        int       `json:"id"`
    PostID    int       `json:"post_id"`
    UserID    int       `json:"user_id"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    User      *User     `json:"user,omitempty"`
    Post      *Post     `json:"post,omitempty"`
}

func (c *Comment) Validate() error {
    if c.PostID == 0 {
        return errors.New("post_id is required")
    }
    if c.UserID == 0 {
        return errors.New("user_id is required")
    }
    if len(c.Content) == 0 {
        return errors.New("content cannot be empty")
    }
    if len(c.Content) > 140 {
        return errors.New("content cannot exceed 140 characters")
    }
    return nil
}