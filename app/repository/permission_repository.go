package repository

import (
    "context"
    "uas-backend-go/app/model"
    "github.com/jackc/pgx/v5/pgxpool"
)

type PermissionRepository struct {
    DB *pgxpool.Pool
}

func NewPermissionRepository(db *pgxpool.Pool) *PermissionRepository {
    return &PermissionRepository{DB: db}
}

func (r *PermissionRepository) GetByRoleID(ctx context.Context, roleID string) ([]model.Permission, error) {

    rows, err := r.DB.Query(ctx, `
        SELECT 
            p.id, 
            p.name, 
            p.resource,
            p.action,
            p.description, 
            p.created_at
        FROM permissions p
        JOIN role_permissions rp ON p.id = rp.permission_id
        WHERE rp.role_id = $1
    `, roleID)

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var perms []model.Permission

    for rows.Next() {
        var perm model.Permission

        err := rows.Scan(
            &perm.ID,
            &perm.Name,
            &perm.Resource,
            &perm.Action,
            &perm.Description,
            &perm.CreatedAt,
        )
        if err != nil {
            return nil, err
        }

        perms = append(perms, perm)
    }

    return perms, nil
}
