package repository

import (
    "context"
    "uas-backend-go/app/model"
    "github.com/jackc/pgx/v5"
)

type PermissionRepository struct {
    DB *pgx.Conn
}

func NewPermissionRepository(db *pgx.Conn) *PermissionRepository {
    return &PermissionRepository{DB: db}
}

func (r *PermissionRepository) GetByRoleID(ctx context.Context, roleID string) ([]model.Permission, error) {
    rows, err := r.DB.Query(ctx,
        `SELECT p.id, p.name, p.description, p.created_at
        FROM permissions p
        JOIN role_permissions rp ON p.id = rp.permission_id
        WHERE rp.role_id=$1`,
        roleID,
    )
    if err != nil {
        return nil, err
    }

    var perms []model.Permission
    for rows.Next() {
        var perm model.Permission
        rows.Scan(&perm.ID, &perm.Name, &perm.Description, &perm.CreatedAt)
        perms = append(perms, perm)
    }

    return perms, nil
}
