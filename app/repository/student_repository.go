package repository

import (
    "context"
    "uas-backend-go/app/model"

    "github.com/jackc/pgx/v5/pgxpool"
)

type StudentRepository struct {
    DB *pgxpool.Pool
}

func NewStudentRepository(db *pgxpool.Pool) *StudentRepository {
    return &StudentRepository{DB: db}
}

func (r *StudentRepository) GetAll(ctx context.Context) ([]model.Student, error) {
    rows, err := r.DB.Query(ctx,
        `SELECT id, user_id, student_id, program_study, academic_year, advisor_id, created_at 
         FROM students ORDER BY created_at DESC`,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []model.Student
    for rows.Next() {
        var s model.Student
        err := rows.Scan(
            &s.ID, &s.UserID, &s.StudentID, &s.ProgramStudy,
            &s.AcademicYear, &s.AdvisorID, &s.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        list = append(list, s)
    }

    return list, nil
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

func (r *StudentRepository) UpdateAdvisor(ctx context.Context, id string, advisorID *string) error {
    _, err := r.DB.Exec(ctx,
        `UPDATE students SET advisor_id=$1 WHERE id=$2`,
        advisorID, id,
    )
    return err
}

func (r *StudentRepository) GetByAdvisor(ctx context.Context, advisorID string) ([]model.Student, error) {
    rows, err := r.DB.Query(ctx,
        `SELECT id, user_id, student_id, program_study, academic_year, advisor_id, created_at
         FROM students WHERE advisor_id = $1`, advisorID,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []model.Student
    for rows.Next() {
        var s model.Student
        err := rows.Scan(
            &s.ID, &s.UserID, &s.StudentID, &s.ProgramStudy,
            &s.AcademicYear, &s.AdvisorID, &s.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        list = append(list, s)
    }

    return list, nil
}
