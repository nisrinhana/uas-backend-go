package model

type GlobalStatistics struct {
    TotalStudents     int `json:"total_students"`
    TotalLecturers    int `json:"total_lecturers"`
    TotalAchievements int `json:"total_achievements"`

    DraftCount     int `json:"draft"`
    SubmittedCount int `json:"submitted"`
    VerifiedCount  int `json:"verified"`
    RejectedCount  int `json:"rejected"`
}

type StudentStatistics struct {
    StudentID        string  `json:"student_id"`
    Total            int     `json:"total"`
    DraftCount       int     `json:"draft"`
    SubmittedCount   int     `json:"submitted"`
    VerifiedCount    int     `json:"verified"`
    RejectedCount    int     `json:"rejected"`
    ProgressPercent  float64 `json:"progress_percent,omitempty"`
}
