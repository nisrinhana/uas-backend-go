package model

import "time"

// Student represents a student entity.
// @swagger:model Student
type Student struct {
    ID           string    `db:"id"`
    UserID       string    `db:"user_id"`
    StudentID    string    `db:"student_id"`
    ProgramStudy string    `db:"program_study"`
    AcademicYear string    `db:"academic_year"`
    AdvisorID    *string   `db:"advisor_id"` 
    CreatedAt    time.Time `db:"created_at"`
}
