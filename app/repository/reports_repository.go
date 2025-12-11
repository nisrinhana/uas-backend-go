package repository

import (
    "context"
    "uas-backend-go/app/model"

    "github.com/jackc/pgx/v5/pgxpool"
)

type ReportRepository struct {
    DB *pgxpool.Pool
}

func NewReportRepository(db *pgxpool.Pool) *ReportRepository {
    return &ReportRepository{DB: db}
}

func (r *ReportRepository) GetGlobalStatistics(ctx context.Context) (model.GlobalStatistics, error) {
    var s model.GlobalStatistics

    queries := []string{
        "SELECT COUNT(*) FROM students",
        "SELECT COUNT(*) FROM lecturers",
        "SELECT COUNT(*) FROM achievement_references",
        "SELECT COUNT(*) FROM achievement_references WHERE status='draft'",
        "SELECT COUNT(*) FROM achievement_references WHERE status='submitted'",
        "SELECT COUNT(*) FROM achievement_references WHERE status='verified'",
        "SELECT COUNT(*) FROM achievement_references WHERE status='rejected'",
    }

    results := []*int{
        &s.TotalStudents,
        &s.TotalLecturers,
        &s.TotalAchievements,
        &s.DraftCount,
        &s.SubmittedCount,
        &s.VerifiedCount,
        &s.RejectedCount,
    }

    for i, q := range queries {
        if err := r.DB.QueryRow(ctx, q).Scan(results[i]); err != nil {
            return s, err
        }
    }

    return s, nil
}

func (r *ReportRepository) GetStudentStatistics(ctx context.Context, studentID string) (model.StudentStatistics, error) {
    var s model.StudentStatistics
    s.StudentID = studentID

    queries := []string{
        "SELECT COUNT(*) FROM achievement_references WHERE student_id=$1",
        "SELECT COUNT(*) FROM achievement_references WHERE student_id=$1 AND status='draft'",
        "SELECT COUNT(*) FROM achievement_references WHERE student_id=$1 AND status='submitted'",
        "SELECT COUNT(*) FROM achievement_references WHERE student_id=$1 AND status='verified'",
        "SELECT COUNT(*) FROM achievement_references WHERE student_id=$1 AND status='rejected'",
    }

    results := []*int{
        &s.Total, &s.DraftCount, &s.SubmittedCount, &s.VerifiedCount, &s.RejectedCount,
    }

    for i, q := range queries {
        if err := r.DB.QueryRow(ctx, q, studentID).Scan(results[i]); err != nil {
            return s, err
        }
    }

    if s.Total > 0 {
        s.ProgressPercent = float64(s.VerifiedCount) / float64(s.Total) * 100
    }

    return s, nil
}
