package service

import (
    "context"
    "uas-backend-go/app/model"
    "uas-backend-go/app/repository"
)

type ReportService struct {
    Repo *repository.ReportRepository
}

func NewReportService(r *repository.ReportRepository) *ReportService {
    return &ReportService{Repo: r}
}

func (s *ReportService) GetGlobalStatistics(ctx context.Context) (model.GlobalStatistics, error) {
    return s.Repo.GetGlobalStatistics(ctx)
}

func (s *ReportService) GetStudentStatistics(ctx context.Context, studentID string) (model.StudentStatistics, error) {
    return s.Repo.GetStudentStatistics(ctx, studentID)
}
