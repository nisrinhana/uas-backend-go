package repository

import (
    "context"
    "uas-backend-go/app/model"
    "github.com/jackc/pgx/v5"
)

type LecturerRepository struct {
    DB *pgx.Conn
}

func NewLecturerRepository(db *pgx.Conn) *LecturerRepository {
    return &LecturerRepository{DB: db}
}

func (r *LecturerRepository) GetByID(ctx context.Context, id string) (model.Lecturer, error) {
    var l model.Lecturer
    err := r.DB.QueryRow(ctx,
        `SELECT id, user_id, lecturer_id, department, created_at
        FROM lecturers WHERE id=$1`, id,
    ).Scan(&l.ID, &l.UserID, &l.LecturerID, &l.Department, &l.CreatedAt)

    return l, err
}
