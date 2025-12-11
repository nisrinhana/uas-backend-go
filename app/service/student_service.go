package service

import (
    "context"
    "uas-backend-go/app/model"
    "uas-backend-go/app/repository"
)

type StudentService struct {
    StudentRepo *repository.StudentRepository
    AchRepo     *repository.AchievementRefRepository
}

func NewStudentService(studentRepo *repository.StudentRepository, achRefRepo *repository.AchievementRefRepository) *StudentService {
    return &StudentService{
        StudentRepo: studentRepo,
        AchRepo:     achRefRepo,
    }
}

func (s *StudentService) GetAll(ctx context.Context) ([]model.Student, error) {
    return s.StudentRepo.GetAll(ctx)
}

func (s *StudentService) GetByID(ctx context.Context, id string) (model.Student, error) {
    return s.StudentRepo.GetByID(ctx, id)
}

func (s *StudentService) GetAchievements(ctx context.Context, id string) ([]model.AchievementReference, error) {
    return s.AchRepo.GetByStudentID(ctx, id)
}

func (s *StudentService) UpdateAdvisor(ctx context.Context, id string, advisorID *string) error {
    return s.StudentRepo.UpdateAdvisor(ctx, id, advisorID)
}
