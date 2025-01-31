// internal/services/user_service.go
package services

import (
	"context"
	"twitter-clone-api/internal/database"
	"twitter-clone-api/internal/models"
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
    return s.repo.Create(ctx, user)
}

func (s *UserService) GetByID(ctx context.Context, id string) (*models.User, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *UserService) GetAll(ctx context.Context) ([]models.User, error) {
    return s.repo.GetAll(ctx)
}

func (s *UserService) Update(ctx context.Context, id string, user *models.User) error {
    return s.repo.Update(ctx, id, user)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}