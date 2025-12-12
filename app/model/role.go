package model

import "time"


// Role defines user roles such as admin, student, lecturer.
// @swagger:model Role
type Role struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}
