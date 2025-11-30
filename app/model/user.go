package model

import "time"

type User struct {
    ID           string    `json:"id"`
    Username     string    `json:"username"`
    Email        string    `json:"email"`
    
    // Field ini menerima "password" dari JSON
    PasswordHash string    `json:"password"`

    FullName     string    `json:"fullname"`
    RoleID       string    `json:"role_id"`
    IsActive     bool      `json:"is_active"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}
