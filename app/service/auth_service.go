package service

import (
    "context"
    "errors"

    "golang.org/x/crypto/bcrypt"
    "uas-backend-go/app/model"
    "uas-backend-go/app/repository"
    "uas-backend-go/utils"
)

type AuthService struct {
    UserRepo       *repository.UserRepository
    PermissionRepo *repository.PermissionRepository
}

func NewAuthService(u *repository.UserRepository, p *repository.PermissionRepository) *AuthService {
    return &AuthService{
        UserRepo:       u,
        PermissionRepo: p,
    }
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, model.User, []string, error) {

    user, err := s.UserRepo.GetByUsername(ctx, username)
    if err != nil {
        return "", model.User{}, nil, errors.New("invalid username or password")
    }

    if !user.IsActive {
        return "", model.User{}, nil, errors.New("account inactive")
    }

    if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
        return "", model.User{}, nil, errors.New("invalid username or password")
    }

    roleID := user.RoleID

    perms, err := s.PermissionRepo.GetByRoleID(ctx, roleID)
    if err != nil {
        return "", model.User{}, nil, errors.New("failed load permissions: " + err.Error())
    }

    permNames := []string{}
    for _, p := range perms {
        permNames = append(permNames, p.Name)
    }

    token, err := utils.GenerateJWT(user.ID, user.RoleID, permNames)
    if err != nil {
        return "", model.User{}, nil, err
    }

    return token, user, permNames, nil
}

func (s *AuthService) GetProfile(ctx context.Context, userID string) (model.User, error) {
    return s.UserRepo.GetByID(ctx, userID)
}
