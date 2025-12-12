package model

import "time"

// Lecturer represents a lecturer who advises students.
// @swagger:model Lecturer
type Lecturer struct {
    ID         string    `db:"id"`
    UserID     string    `db:"user_id"`
    LecturerID string    `db:"lecturer_id"`
    Department string    `db:"department"`
    CreatedAt  time.Time `db:"created_at"`
}
