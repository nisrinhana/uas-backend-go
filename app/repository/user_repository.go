package repository

import (
    "context"
    "uas-backend-go/app/model"

    "github.com/jackc/pgx/v5"
)

type UserRepository struct {
    DB *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (model.User, error) {
    var u model.User
    err := r.DB.QueryRow(ctx,
        `SELECT id, username, email, password_hash, full_name,
        role_id, is_active, created_at, updated_at
        FROM users WHERE username=$1`,
        username,
    ).Scan(
        &u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.FullName,
        &u.RoleID, &u.IsActive, &u.CreatedAt, &u.UpdatedAt,
    )

    return u, err
}
