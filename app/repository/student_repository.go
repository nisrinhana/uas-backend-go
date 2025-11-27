package repository

import (
    "context"
    "uas-backend-go/app/model"
    "github.com/jackc/pgx/v5"
)

type StudentRepository struct {
    DB *pgx.Conn
}

func NewStudentRepository(db *pgx.Conn) *StudentRepository {
    return &StudentRepository{DB: db}
}

func (r *StudentRepository) GetByID(ctx context.Context, id string) (model.Student, error) {
    var s model.Student
    err := r.DB.QueryRow(ctx,
        `SELECT id, user_id, student_id, program_study,
        academic_year, advisor_id, created_at
        FROM students WHERE id=$1`, id,
    ).Scan(
        &s.ID, &s.UserID, &s.StudentID, &s.ProgramStudy,
        &s.AcademicYear, &s.AdvisorID, &s.CreatedAt,
    )

    return s, err
}
