package repository

import (
    "context"
    "errors"
    "uas-backend-go/app/model"

    "github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
    DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
    return &UserRepository{DB: db}
}

// GetAll mengambil semua user
func (r *UserRepository) GetAll(ctx context.Context) ([]model.User, error) {
    rows, err := r.DB.Query(ctx, `
        SELECT id, username, email, password_hash, full_name, role_id, is_active, created_at, updated_at
        FROM users
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []model.User
    for rows.Next() {
        var u model.User
        err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.FullName,
            &u.RoleID, &u.IsActive, &u.CreatedAt, &u.UpdatedAt)
        if err != nil {
            return nil, err
        }
        users = append(users, u)
    }

    return users, nil
}

// GetByID mengambil user berdasarkan ID
func (r *UserRepository) GetByID(ctx context.Context, id string) (model.User, error) {
    var u model.User
    err := r.DB.QueryRow(ctx, `
        SELECT id, username, email, password_hash, full_name, role_id, is_active, created_at, updated_at
        FROM users
        WHERE id = $1
    `, id).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.FullName,
        &u.RoleID, &u.IsActive, &u.CreatedAt, &u.UpdatedAt)
    if err != nil {
        return model.User{}, err
    }
    return u, nil
}

// Create menambahkan user baru
func (r *UserRepository) Create(ctx context.Context, user model.User) error {
    commandTag, err := r.DB.Exec(ctx, `
        INSERT INTO users (id, username, email, password_hash, full_name, role_id, is_active, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7,NOW(),NOW())
    `,
        user.ID, user.Username, user.Email, user.PasswordHash, user.FullName,
        user.RoleID, user.IsActive,
    )
    if err != nil {
        return err
    }
    if commandTag.RowsAffected() != 1 {
        return errors.New("failed to insert user")
    }
    return nil
}

// Update mengubah data user
func (r *UserRepository) Update(ctx context.Context, id string, user model.User) error {
    commandTag, err := r.DB.Exec(ctx, `
        UPDATE users
        SET username=$1, email=$2, password_hash=$3, full_name=$4, role_id=$5, is_active=$6, updated_at=NOW()
        WHERE id=$7
    `,
        user.Username, user.Email, user.PasswordHash, user.FullName, user.RoleID, user.IsActive, id,
    )
    if err != nil {
        return err
    }
    if commandTag.RowsAffected() != 1 {
        return errors.New("user not found or no changes made")
    }
    return nil
}

// Delete menghapus user berdasarkan ID
func (r *UserRepository) Delete(ctx context.Context, id string) error {
    commandTag, err := r.DB.Exec(ctx, `
        DELETE FROM users WHERE id=$1
    `, id)
    if err != nil {
        return err
    }
    if commandTag.RowsAffected() != 1 {
        return errors.New("user not found")
    }
    return nil
}

// GetByUsername mengambil user berdasarkan username
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (model.User, error) {
    var u model.User
    err := r.DB.QueryRow(ctx, `
        SELECT id, username, email, password_hash, full_name, role_id, is_active, created_at, updated_at
        FROM users
        WHERE username = $1
    `, username).Scan(
        &u.ID, &u.Username, &u.Email, &u.PasswordHash, &u.FullName,
        &u.RoleID, &u.IsActive, &u.CreatedAt, &u.UpdatedAt,
    )
    if err != nil {
        return model.User{}, err
    }
    return u, nil
}
