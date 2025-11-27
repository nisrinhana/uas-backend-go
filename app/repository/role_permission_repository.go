package repository

import (
    "context"
    "uas-backend-go/app/model"
    "github.com/jackc/pgx/v5"
)

type RolePermissionRepository struct {
    DB *pgx.Conn
}

func NewRolePermissionRepository(db *pgx.Conn) *RolePermissionRepository {
    return &RolePermissionRepository{DB: db}
}

func (r *RolePermissionRepository) AssignPermission(ctx context.Context, rp model.RolePermission) error {
    _, err := r.DB.Exec(ctx,
        `INSERT INTO role_permissions (role_id, permission_id)
        VALUES ($1, $2)`,
        rp.RoleID, rp.PermissionID,
    )
    return err
}
