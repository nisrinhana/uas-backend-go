package helper

import (
    "context"
    "net/http"

    "github.com/google/uuid"
    "uas-backend-go/app/model"

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


// GetAll godoc
// @Summary Get all users
// @Description Mengambil seluruh data user
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {array} model.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (h *UserHelper) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := h.Service.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}


// GetByID godoc
// @Summary Get user by ID
// @Description Mengambil detail user berdasarkan ID
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
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


// Create godoc
// @Summary Create a new user
// @Description Membuat user baru
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user body model.User true "User Data"
// @Success 201 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *UserHelper) Create(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    user.ID = uuid.New().String()

    if user.Email == "" {
        user.Email = user.Username + "@example.com"
    }

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


// Update godoc
// @Summary Update user
// @Description Mengupdate data user
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body model.User true "Updated User Data"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [put]
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


// Delete godoc
// @Summary Delete user
// @Description Menghapus user berdasarkan ID
// @Tags Users
// @Security BearerAuth
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [delete]
func (h *UserHelper) Delete(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	if err := h.Service.Delete(ctx, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}


// UpdateRole godoc
// @Summary Update user's role
// @Description Mengganti role user
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param role body map[string]string true "Role Payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id}/role [put]
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
