package repository

import (
    "context"
    "errors"
    "uas-backend-go/app/model"

    "github.com/jackc/pgx/v5/pgxpool"
)

type AchievementRefRepository struct {
    DB *pgxpool.Pool
}

func NewAchievementRefRepository(db *pgxpool.Pool) *AchievementRefRepository {
    return &AchievementRefRepository{DB: db}
}

//
// CREATE REFERENCE
//
func (r *AchievementRefRepository) Create(ctx context.Context, ref model.AchievementReference) error {
    query := `
        INSERT INTO achievement_references 
        (id, student_id, mongo_achievement_id, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())
    `

    _, err := r.DB.Exec(ctx, query,
        ref.ID,
        ref.StudentID,
        ref.MongoAchievementID,
        ref.Status,
    )

    return err
}

//
// UPDATE STATUS (submit, verify, reject)
//
func (r *AchievementRefRepository) UpdateStatus(ctx context.Context, ref model.AchievementReference) error {
    query := `
        UPDATE achievement_references
        SET status = $1,
            submitted_at = $2,
            verified_at = $3,
            verified_by = $4,
            rejection_note = $5,
            updated_at = NOW()
        WHERE id = $6
    `

    cmd, err := r.DB.Exec(ctx, query,
        ref.Status,
        ref.SubmittedAt,
        ref.VerifiedAt,
        ref.VerifiedBy,
        ref.RejectionNote,
        ref.ID,
    )

    if err != nil {
        return err
    }

    if cmd.RowsAffected() == 0 {
        return errors.New("achievement reference not found")
    }

    return nil
}

func (r *AchievementRefRepository) SoftDelete(ctx context.Context, id string) error {
    query := `
        UPDATE achievement_references
        SET deleted_at = NOW(), 
            status = 'deleted',
            updated_at = NOW()
        WHERE id = $1
    `

    cmd, err := r.DB.Exec(ctx, query, id)
    if err != nil {
        return err
    }

    if cmd.RowsAffected() == 0 {
        return errors.New("achievement reference not found")
    }

    return nil
}

//
// GET BY ID
//
func (r *AchievementRefRepository) GetByID(ctx context.Context, id string) (model.AchievementReference, error) {
    var ref model.AchievementReference

    query := `
        SELECT id, student_id, mongo_achievement_id, status,
               submitted_at, verified_at, verified_by, rejection_note,
               deleted_at, created_at, updated_at
        FROM achievement_references
        WHERE id = $1 AND deleted_at IS NULL
    `

    err := r.DB.QueryRow(ctx, query, id).Scan(
        &ref.ID,
        &ref.StudentID,
        &ref.MongoAchievementID,
        &ref.Status,
        &ref.SubmittedAt,
        &ref.VerifiedAt,
        &ref.VerifiedBy,
        &ref.RejectionNote,
        &ref.DeletedAt,
        &ref.CreatedAt,
        &ref.UpdatedAt,
    )

    return ref, err
}

//
// GET BY STUDENT ID (mahasiswa lihat own achievements)
//
func (r *AchievementRefRepository) GetByStudentID(ctx context.Context, studentID string) ([]model.AchievementReference, error) {
    query := `
        SELECT id, student_id, mongo_achievement_id, status,
               submitted_at, verified_at, verified_by, rejection_note,
               deleted_at, created_at, updated_at
        FROM achievement_references
        WHERE student_id = $1 AND deleted_at IS NULL
        ORDER BY created_at DESC
    `

    rows, err := r.DB.Query(ctx, query, studentID)
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var list []model.AchievementReference

    for rows.Next() {
        var ref model.AchievementReference
        err := rows.Scan(
            &ref.ID,
            &ref.StudentID,
            &ref.MongoAchievementID,
            &ref.Status,
            &ref.SubmittedAt,
            &ref.VerifiedAt,
            &ref.VerifiedBy,
            &ref.RejectionNote,
            &ref.DeletedAt,
            &ref.CreatedAt,
            &ref.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        list = append(list, ref)
    }

    return list, nil
}

//
// GET BY ADVISOR (dosen wali)
//
func (r *AchievementRefRepository) GetByAdvisor(ctx context.Context, advisorID string) ([]model.AchievementReference, error) {
    query := `
        SELECT ar.id, ar.student_id, ar.mongo_achievement_id, ar.status,
               ar.submitted_at, ar.verified_at, ar.verified_by, ar.rejection_note,
               ar.deleted_at, ar.created_at, ar.updated_at
        FROM achievement_references ar
        JOIN students s ON s.id = ar.student_id
        WHERE s.advisor_id = $1 AND ar.deleted_at IS NULL
        ORDER BY ar.created_at DESC
    `

    rows, err := r.DB.Query(ctx, query, advisorID)
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var list []model.AchievementReference

    for rows.Next() {
        var ref model.AchievementReference
        err := rows.Scan(
            &ref.ID,
            &ref.StudentID,
            &ref.MongoAchievementID,
            &ref.Status,
            &ref.SubmittedAt,
            &ref.VerifiedAt,
            &ref.VerifiedBy,
            &ref.RejectionNote,
            &ref.DeletedAt,
            &ref.CreatedAt,
            &ref.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        list = append(list, ref)
    }

    return list, nil
}

func (r *AchievementRefRepository) GetAll(ctx context.Context) ([]model.AchievementReference, error) {
    query := `
    SELECT 
        ar.id,
        ar.student_id,
        ar.mongo_achievement_id,
        ar.status,
        ar.created_at,

        s.user_id,
        s.program_study,
        s.academic_year,
        u.username
    FROM achievement_references ar
    LEFT JOIN students s ON ar.student_id = s.id
    LEFT JOIN users u ON s.user_id = u.id
    WHERE ar.deleted_at IS NULL
    ORDER BY ar.created_at DESC
    `

    rows, err := r.DB.Query(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var refs []model.AchievementReference

    for rows.Next() {
        var ref model.AchievementReference

        err := rows.Scan(
            &ref.ID,
            &ref.StudentID,
            &ref.MongoAchievementID,
            &ref.Status,
            &ref.CreatedAt,

            &ref.UserID,
            &ref.ProgramStudy,
            &ref.AcademicYear,
            &ref.Username,
        )

        if err != nil {
            return nil, err
        }

        refs = append(refs, ref)
    }

    return refs, nil
}

func (r *AchievementRefRepository) GetByMongoID(ctx context.Context, mongoID string) (model.AchievementReference, error) {
    var ref model.AchievementReference

    query := `
        SELECT id, student_id, mongo_achievement_id, status,
               submitted_at, verified_at, verified_by,
               rejection_note, deleted_at, created_at, updated_at
        FROM achievement_references
        WHERE mongo_achievement_id = $1 AND deleted_at IS NULL
        LIMIT 1;
    `
    err := r.DB.QueryRow(ctx, query, mongoID).Scan(
        &ref.ID, &ref.StudentID, &ref.MongoAchievementID, &ref.Status,
        &ref.SubmittedAt, &ref.VerifiedAt, &ref.VerifiedBy,
        &ref.RejectionNote, &ref.DeletedAt, &ref.CreatedAt, &ref.UpdatedAt,
    )

    return ref, err
}
