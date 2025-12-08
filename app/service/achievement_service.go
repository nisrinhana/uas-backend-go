package service

import (
    "context"
    "time"
    "uas-backend-go/app/model"
    "uas-backend-go/app/repository"

    "github.com/google/uuid"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type AchievementService struct {
    RefRepo   *repository.AchievementRefRepository
    MongoRepo *repository.AchievementMongoRepository
}

func NewAchievementService(ref *repository.AchievementRefRepository, mongo *repository.AchievementMongoRepository) *AchievementService {
    return &AchievementService{RefRepo: ref, MongoRepo: mongo}
}

//
// CREATE DRAFT
//
func (s *AchievementService) Create(ctx context.Context, studentID string, m model.Achievement) (string, error) {

    now := time.Now()

    m.Status = "draft"
    m.CreatedAt = now
    m.UpdatedAt = now
    m.VerifiedAt = nil
    m.VerifiedBy = nil
    m.RejectionNote = nil
    m.StudentID = studentID

    mongoID, err := s.MongoRepo.Create(ctx, m)
    if err != nil {
        return "", err
    }

    ref := model.AchievementReference{
        ID:                 uuid.New().String(),
        StudentID:          studentID,
        MongoAchievementID: mongoID.Hex(),
        Status:             "draft",
        CreatedAt:          now,
        UpdatedAt:          now,
    }

    err = s.RefRepo.Create(ctx, ref)
    return ref.ID, err
}

//
// UPDATE DRAFT ONLY
//
func (s *AchievementService) Update(ctx context.Context, mongoID string, m model.Achievement) error {

    id, _ := primitive.ObjectIDFromHex(mongoID)

    // updatedAt harus diperbarui
    m.UpdatedAt = time.Now()

    return s.MongoRepo.Update(ctx, id, m)
}


// DELETE (draft only)
func (s *AchievementService) Delete(ctx context.Context, refID string, mongoID string) (model.Achievement, error) {
    objID, _ := primitive.ObjectIDFromHex(mongoID)

    // Snapshot sebelum delete
    snapshot, err := s.MongoRepo.GetByID(ctx, objID)
    if err != nil {
        return model.Achievement{}, err
    }

    // Delete Mongo
    err = s.MongoRepo.Delete(ctx, objID)
    if err != nil {
        return model.Achievement{}, err
    }

    // Soft delete + status=deleted
    err = s.RefRepo.SoftDelete(ctx, refID)
    if err != nil {
        return model.Achievement{}, err
    }

    return snapshot, nil
}

//
// SUBMIT
//
func (s *AchievementService) Submit(ctx context.Context, ref model.AchievementReference) error {
    now := time.Now()
    ref.SubmittedAt = &now
    ref.Status = "submitted"

    // update status di Mongo
    mongoID, _ := primitive.ObjectIDFromHex(ref.MongoAchievementID)
    s.MongoRepo.Update(ctx, mongoID, model.Achievement{
        Status:    "submitted",
        UpdatedAt: now,
    })

    return s.RefRepo.UpdateStatus(ctx, ref)
}

//
// VERIFY
//
func (s *AchievementService) Verify(ctx context.Context, ref model.AchievementReference, lecturerID string) error {

    now := time.Now()
    ref.VerifiedAt = &now
    ref.VerifiedBy = &lecturerID
    ref.Status = "verified"

    mongoID, _ := primitive.ObjectIDFromHex(ref.MongoAchievementID)
    s.MongoRepo.Update(ctx, mongoID, model.Achievement{
        Status:     "verified",
        VerifiedAt: &now,
        VerifiedBy: &lecturerID,
        UpdatedAt:  now,
    })

    return s.RefRepo.UpdateStatus(ctx, ref)
}

//
// REJECT
//
func (s *AchievementService) Reject(ctx context.Context, ref model.AchievementReference, note string) error {
    ref.RejectionNote = &note
    ref.Status = "rejected"

    now := time.Now()

    mongoID, _ := primitive.ObjectIDFromHex(ref.MongoAchievementID)
    s.MongoRepo.Update(ctx, mongoID, model.Achievement{
        Status:        "rejected",
        RejectionNote: &note,
        UpdatedAt:     now,
    })

    return s.RefRepo.UpdateStatus(ctx, ref)
}

//
// GET FUNCTIONS
//
func (s *AchievementService) GetByID(ctx context.Context, id primitive.ObjectID) (model.Achievement, error) {
    return s.MongoRepo.GetByID(ctx, id)
}

func (s *AchievementService) GetRefByID(ctx context.Context, id string) (model.AchievementReference, error) {
    return s.RefRepo.GetByID(ctx, id)
}

func (s *AchievementService) GetForStudent(ctx context.Context, studentID string) ([]model.AchievementReference, error) {
    return s.RefRepo.GetByStudentID(ctx, studentID)
}

func (s *AchievementService) GetForAdvisor(ctx context.Context, advisorID string) ([]model.AchievementReference, error) {
    return s.RefRepo.GetByAdvisor(ctx, advisorID)
}

func (s *AchievementService) GetAll(ctx context.Context) ([]model.AchievementReference, error) {
    return s.RefRepo.GetAll(ctx)
}

func (s *AchievementService) GetRefByMongoID(ctx context.Context, mongoID string) (model.AchievementReference, error) {
    return s.RefRepo.GetByMongoID(ctx, mongoID)
}
