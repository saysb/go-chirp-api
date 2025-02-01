package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"twitter-clone-api/internal/models"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{
        db: db,
    }
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
    query := `
        INSERT INTO users (username, email, password, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW())
        RETURNING id, created_at, updated_at`

    err := r.db.QueryRowContext(ctx, query,
        user.Username,
        user.Email,
        user.Password,
    ).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

    if err != nil {
        return err
    }

    return nil
}

func (r *UserRepository) GetByEmailOrUsername(ctx context.Context, email string, username string) (*models.User, error) {
    user := &models.User{}
    query := `SELECT id, username, email, created_at, updated_at FROM users WHERE email = $1 OR username = $2`

    err := r.db.QueryRowContext(ctx, query, email, username).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.CreatedAt,
        &user.UpdatedAt,
    )

    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }

    return user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
    user := &models.User{}
    query := `
        SELECT id, username, email, created_at, updated_at
        FROM users
        WHERE id = $1`

    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.CreatedAt,
        &user.UpdatedAt,
    )

    if err == sql.ErrNoRows {
        return nil, errors.New("user not found")
    }
    if err != nil {
        return nil, err
    }

    return user, nil
}


func (r *UserRepository) Update(ctx context.Context, id string, user *models.User) error {
    query := `
        UPDATE users
        SET username = $1, email = $2, updated_at = NOW()
        WHERE id = $3
        RETURNING created_at, updated_at`

    err := r.db.QueryRowContext(ctx, query,
        user.Username,
        user.Email,
        id,
    ).Scan(&user.CreatedAt, &user.UpdatedAt)

    if err == sql.ErrNoRows {
        return errors.New("user not found")
    }
    if err != nil {
        return err
    }

    user.ID, err = strconv.Atoi(id)
    if err != nil {
        return fmt.Errorf("invalid user ID: %v", err)
    }
    return nil
}