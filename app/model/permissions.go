package model

import "time"


// Permission defines access control capabilities.
// @swagger:model Permission
type Permission struct {
    ID          string    `db:"id"`
    Name        string    `db:"name"`
    Resource    string    `db:"resource"`
    Action      string    `db:"action"`
    Description string    `db:"description"`
    CreatedAt   time.Time `db:"created_at"`
}

