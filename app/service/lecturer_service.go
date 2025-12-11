package service

import (
    "context"
    "uas-backend-go/app/model"
    "uas-backend-go/app/repository"
)

type LecturerService struct {
    LecturerRepo *repository.LecturerRepository
    StudentRepo  *repository.StudentRepository
}

func NewLecturerService(lr *repository.LecturerRepository, sr *repository.StudentRepository) *LecturerService {
    return &LecturerService{LecturerRepo: lr, StudentRepo: sr}
}

func (s *LecturerService) GetAll(ctx context.Context) ([]model.Lecturer, error) {
    return s.LecturerRepo.GetAll(ctx)
}

func (s *LecturerService) GetAdvisees(ctx context.Context, lecturerID string) ([]model.Student, error) {
    return s.StudentRepo.GetByAdvisor(ctx, lecturerID)
}
