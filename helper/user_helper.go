package helper

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"uas-backend-go/app/model"
	"uas-backend-go/utils"

	"github.com/gin-gonic/gin"
)

// Interface untuk UserService agar helper tidak perlu import package service
type UserServiceInterface interface {
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, id string) (model.User, error)
	Create(ctx context.Context, u model.User) error
	Update(ctx context.Context, id string, u model.User) error
	Delete(ctx context.Context, id string) error
	UpdateRole(ctx context.Context, id string, roleID string) error
}

type UserHelper struct {
	Service UserServiceInterface
}


func NewUserHelper(s UserServiceInterface) *UserHelper {
	return &UserHelper{Service: s}
}

// GET /api/v1/users
func (h *UserHelper) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := h.Service.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GET /api/v1/users/:id
func (h *UserHelper) GetByID(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	user, err := h.Service.GetByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// POST /api/v1/users
func (h *UserHelper) Create(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// ID baru
	user.ID = uuid.New().String()
	
// Hash password langsung setelah binding
if user.PasswordHash != "" {
user.PasswordHash = utils.HashPassword(user.PasswordHash)

}


	// Email default kalau kosong (optional)
	if user.Email == "" {
		user.Email = user.Username + "@example.com"
	}

	// Pastikan RoleID valid UUID
	if _, err := uuid.Parse(user.RoleID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role_id"})
		return
	}

	ctx := c.Request.Context()
	if err := h.Service.Create(ctx, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// PUT /api/v1/users/:id
func (h *UserHelper) Update(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	ctx := c.Request.Context()
	if err := h.Service.Update(ctx, id, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DELETE /api/v1/users/:id
func (h *UserHelper) Delete(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	if err := h.Service.Delete(ctx, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

// PUT /api/v1/users/:id/role
func (h *UserHelper) UpdateRole(c *gin.Context) {
    id := c.Param("id")

    var payload struct {
        RoleID string `json:"role_id"`
    }

    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    ctx := c.Request.Context()

    if err := h.Service.UpdateRole(ctx, id, payload.RoleID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "role updated successfully",
        "user_id": id,
        "new_role": payload.RoleID,
    })
}
