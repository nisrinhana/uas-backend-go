package model


// RolePermission defines many-to-many mapping of roles and permissions.
// @swagger:model RolePermission
type RolePermission struct {
    RoleID       string `db:"role_id"`
    PermissionID string `db:"permission_id"`
}
