package repository

import (
    "context"
    "uas-backend-go/app/model"
    "github.com/jackc/pgx/v5"
)

type AchievementRefRepository struct {
    DB *pgx.Conn
}

func NewAchievementRefRepository(db *pgx.Conn) *AchievementRefRepository {
    return &AchievementRefRepository{DB: db}
}

func (r *AchievementRefRepository) Create(ctx context.Context, ref model.AchievementReference) error {
    _, err := r.DB.Exec(ctx,
        `INSERT INTO achievement_references
        (id, student_id, mongo_achievement_id, status, created_at, updated_at)
        VALUES ($1, $2, $3, 'draft', NOW(), NOW())`,
        ref.ID, ref.StudentID, ref.MongoAchievementID,
    )
    return err
}

func (r *AchievementRefRepository) UpdateStatus(ctx context.Context, id string, status string) error {
    _, err := r.DB.Exec(ctx,
        `UPDATE achievement_references
        SET status=$1, updated_at=NOW()
        WHERE id=$2`,
        status, id,
    )
    return err
}

func (r *AchievementRefRepository) SetVerification(ctx context.Context, id string, verifiedBy string) error {
    _, err := r.DB.Exec(ctx,
        `UPDATE achievement_references
        SET status='verified', verified_by=$1, verified_at=NOW()
        WHERE id=$2`,
        verifiedBy, id,
    )
    return err
}

func (r *AchievementRefRepository) SetRejection(ctx context.Context, id string, note string) error {
    _, err := r.DB.Exec(ctx,
        `UPDATE achievement_references
        SET status='rejected', rejection_note=$1, updated_at=NOW()
        WHERE id=$2`,
        note, id,
    )
    return err
}

func (r *AchievementRefRepository) GetByStudent(ctx context.Context, studentID string) ([]model.AchievementReference, error) {
    rows, err := r.DB.Query(ctx,
        `SELECT id, student_id, mongo_achievement_id, status,
        submitted_at, verified_at, verified_by, rejection_note,
        created_at, updated_at
        FROM achievement_references
        WHERE student_id=$1`,
        studentID,
    )
    if err != nil {
        return nil, err
    }

    var refs []model.AchievementReference
    for rows.Next() {
        var ref model.AchievementReference
        rows.Scan(
            &ref.ID, &ref.StudentID, &ref.MongoAchievementID, &ref.Status,
            &ref.SubmittedAt, &ref.VerifiedAt, &ref.VerifiedBy, &ref.RejectionNote,
            &ref.CreatedAt, &ref.UpdatedAt,
        )
        refs = append(refs, ref)
    }

    return refs, nil
}
