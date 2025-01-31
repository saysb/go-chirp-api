// internal/models/post.go
package models

import (
	"errors"
	"time"
)

type Post struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    User      *User     `json:"user,omitempty"`  // Relation avec User, omitempty signifie que le champ sera omis si nil
}

func (p *Post) Validate() error {
    if p.UserID == 0 {
        return errors.New("user_id is required")
    }
    if len(p.Content) == 0 {
        return errors.New("content cannot be empty")
    }
    if len(p.Content) > 280 {  // Comme Twitter
        return errors.New("content cannot exceed 280 characters")
    }
    return nil
}