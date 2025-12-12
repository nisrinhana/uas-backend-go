package service

import (
    "context"
    "golang.org/x/crypto/bcrypt"
    "uas-backend-go/app/model"
    "uas-backend-go/app/repository"
)

type UserService struct {
    UserRepo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
    return &UserService{UserRepo: r}
}

func (s *UserService) GetAll(ctx context.Context) ([]model.User, error) {
    return s.UserRepo.GetAll(ctx)
}

func (s *UserService) GetByID(ctx context.Context, id string) (model.User, error) {
    return s.UserRepo.GetByID(ctx, id)
}

func (s *UserService) Create(ctx context.Context, user model.User) error {

    // HASH PASSWORD DULU
    hashed, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.PasswordHash = string(hashed)

    // simpan
    return s.UserRepo.Create(ctx, user)
}

func (s *UserService) Update(ctx context.Context, id string, user model.User) error {

    // update juga harus hashing ulang jika password diubah
    if user.PasswordHash != "" {
        hashed, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
        if err != nil {
            return err
        }
        user.PasswordHash = string(hashed)
    }

    return s.UserRepo.Update(ctx, id, user)
}

func (s *UserService) Delete(ctx context.Context, id string) error {
    return s.UserRepo.Delete(ctx, id)
}

func (s *UserService) UpdateRole(ctx context.Context, id string, roleID string) error {
    return s.UserRepo.UpdateRole(ctx, id, roleID)
}
