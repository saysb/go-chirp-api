// internal/services/user_service.go
package services

import (
	"context"
	"errors"
	"twitter-clone-api/internal/database"
	"twitter-clone-api/internal/models"
	"twitter-clone-api/internal/utils/password"
)

type UserService struct {
    repo *database.UserRepository
}

func NewUserService(repo *database.UserRepository) *UserService {
    return &UserService{
        repo: repo,
    }
}

func (s *UserService) Create(ctx context.Context, user *models.User) error {
    existingUser, err := s.repo.GetByEmailOrUsername(ctx, user.Email, user.Username)
    if err != nil {
        return err
    }
    if existingUser != nil {
        return errors.New("userAlreadyExists")
    }
    newPassword, err := password.HashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = newPassword
    return s.repo.Create(ctx, user)
}

func (s *UserService) GetByID(ctx context.Context, id string) (*models.User, error) {
    return s.repo.GetByID(ctx, id)
}


func (s *UserService) Update(ctx context.Context, id string, user *models.User) error {
    return s.repo.Update(ctx, id, user)
}