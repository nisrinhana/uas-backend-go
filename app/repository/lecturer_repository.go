package repository

import (
    "context"
    "uas-backend-go/app/model"

    "github.com/jackc/pgx/v5/pgxpool"
)

type LecturerRepository struct {
    DB *pgxpool.Pool
}

func NewLecturerRepository(db *pgxpool.Pool) *LecturerRepository {
    return &LecturerRepository{DB: db}
}

func (r *LecturerRepository) GetAll(ctx context.Context) ([]model.Lecturer, error) {
    rows, err := r.DB.Query(ctx,
        `SELECT id, user_id, lecturer_id, department, created_at 
         FROM lecturers ORDER BY created_at DESC`,
    )
    if err != nil { return nil, err }

    defer rows.Close()
    var list []model.Lecturer

    for rows.Next() {
        var l model.Lecturer
        err := rows.Scan(&l.ID, &l.UserID, &l.LecturerID, &l.Department, &l.CreatedAt)
        if err != nil { return nil, err }
        list = append(list, l)
    }

    return list, nil
}

func (r *LecturerRepository) GetByID(ctx context.Context, id string) (model.Lecturer, error) {
    var l model.Lecturer
    err := r.DB.QueryRow(ctx,
        `SELECT id, user_id, lecturer_id, department, created_at
        FROM lecturers WHERE id=$1`, id,
    ).Scan(&l.ID, &l.UserID, &l.LecturerID, &l.Department, &l.CreatedAt)

    return l, err
}
