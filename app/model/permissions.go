package model

import "time"

type Permission struct {
    ID          string    `db:"id"`
    Name        string    `db:"name"`
    Description string    `db:"description"`
    CreatedAt   time.Time `db:"created_at"`
}
