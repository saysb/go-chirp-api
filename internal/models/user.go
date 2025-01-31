// internal/models/user.go
package models

import (
	"errors"
	"time"
)

// User représente un utilisateur dans notre système
// Les tags `json:"..."` définissent comment les champs seront sérialisés en JSON
type User struct {
    ID        int       `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"-"` 
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"` 
}

type UpdateUserRequest struct {
    Username string `json:"username,omitempty"`
    Email    string `json:"email,omitempty"`
}

// Validate vérifie si les données de l'utilisateur sont valides
func (u *User) Validate() error {
    if len(u.Username) < 3 {
        return errors.New("username must be at least 3 characters long")
    }
    if len(u.Email) < 5 {
        return errors.New("invalid email address")
    }
    if len(u.Password) < 6 {
        return errors.New("password must be at least 6 characters long")
    }
    return nil
}