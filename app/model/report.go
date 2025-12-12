package model


// GlobalStatistics represents analytics for the entire system.
// @swagger:model GlobalStatistics
type GlobalStatistics struct {
    TotalStudents     int `json:"total_students"`
    TotalLecturers    int `json:"total_lecturers"`
    TotalAchievements int `json:"total_achievements"`

    DraftCount     int `json:"draft"`
    SubmittedCount int `json:"submitted"`
    VerifiedCount  int `json:"verified"`
    RejectedCount  int `json:"rejected"`
}

// StudentStatistics represents analytics for a single student.
// @swagger:model StudentStatistics
type StudentStatistics struct {
    StudentID        string  `json:"student_id"`
    Total            int     `json:"total"`
    DraftCount       int     `json:"draft"`
    SubmittedCount   int     `json:"submitted"`
    VerifiedCount    int     `json:"verified"`
    RejectedCount    int     `json:"rejected"`
    ProgressPercent  float64 `json:"progress_percent,omitempty"`
}
