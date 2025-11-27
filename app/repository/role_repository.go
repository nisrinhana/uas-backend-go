package repository

import (
    "context"
    "uas-backend-go/app/model"
    "github.com/jackc/pgx/v5"
)

type RoleRepository struct {
    DB *pgx.Conn
}

func NewRoleRepository(db *pgx.Conn) *RoleRepository {
    return &RoleRepository{DB: db}
}

func (r *RoleRepository) GetByID(ctx context.Context, id string) (model.Role, error) {
    var role model.Role
    err := r.DB.QueryRow(ctx,
        `SELECT id, name, description, created_at
        FROM roles WHERE id=$1`, id,
    ).Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt)

    return role, err
}
