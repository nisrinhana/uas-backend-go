package model

import "time"

type AchievementReference struct {
    ID                 string     `db:"id"`
    StudentID          string     `db:"student_id"`
    MongoAchievementID string     `db:"mongo_achievement_id"`
    Status             string     `db:"status"`
    SubmittedAt        *time.Time `db:"submitted_at"`
    VerifiedAt         *time.Time `db:"verified_at"`
    VerifiedBy         *string    `db:"verified_by"`
    RejectionNote      *string    `db:"rejection_note"`
    DeletedAt          *time.Time `db:"deleted_at"`
    CreatedAt          time.Time  `db:"created_at"`
    UpdatedAt          time.Time  `db:"updated_at"`

    // Tambahan untuk JOIN
    UserID       string  `db:"user_id" json:"user_id"`
    Username     string  `db:"username" json:"username"`
    ProgramStudy string  `db:"program_study" json:"program_study"`
    AcademicYear string  `db:"academic_year" json:"academic_year"`
}

