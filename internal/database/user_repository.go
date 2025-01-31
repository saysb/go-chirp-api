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

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
    query := `
        SELECT id, username, email, created_at, updated_at
        FROM users
        ORDER BY id`

    rows, err := r.db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        err := rows.Scan(
            &user.ID,
            &user.Username,
            &user.Email,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
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

func (r *UserRepository) Delete(ctx context.Context, id string) error {
    query := `DELETE FROM users WHERE id = $1`

    result, err := r.db.ExecContext(ctx, query, id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return errors.New("user not found")
    }

    return nil
}