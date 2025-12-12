package model

import "time"

// User represents system user with role-based access.
// @swagger:model User
type User struct {
    ID           string    `json:"id"`
    Username     string    `json:"username"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"password_hash"` 
    FullName     string    `json:"full_name"`
    RoleID       string    `json:"role_id"`
    IsActive     bool      `json:"is_active"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

