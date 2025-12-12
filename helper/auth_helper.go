package helper

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "uas-backend-go/app/model"
    "uas-backend-go/utils"
)

type AuthServiceInterface interface {
    Login(ctx context.Context, username, password string) (string, model.User, []string, error)
    GetProfile(ctx context.Context, userID string) (model.User, error)
}

type AuthHelper struct {
    AuthService AuthServiceInterface
}

func NewAuthHelper(a AuthServiceInterface) *AuthHelper {
    return &AuthHelper{AuthService: a}
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// Login godoc
// @Summary Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (h *AuthHelper) Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    ctx := c.Request.Context()
    token, user, perms, err := h.AuthService.Login(ctx, req.Username, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{
        "token":       token,
        "user":        user,
        "permissions": perms,
    })
}

// Profile godoc
// @Summary Get profile
// @Tags Auth
// @Security BearerAuth
// @Success 200 {object} model.User
// @Failure 401 {object} map[string]string
// @Router /auth/profile [get]
func (h *AuthHelper) Profile(c *gin.Context) {
    userID := c.GetString("user_id")
    if userID == "" {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }

    ctx := c.Request.Context()
    user, err := h.AuthService.GetProfile(ctx, userID)
    if err != nil {
        c.JSON(404, gin.H{"error": "user not found"})
        return
    }

    c.JSON(200, user)
}

// Refresh godoc
// @Summary Refresh JWT token
// @Tags Auth
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Router /auth/refresh [post]
func (h *AuthHelper) Refresh(c *gin.Context) {
    tokenString := c.GetHeader("Authorization")[7:]
    newToken, err := utils.RefreshJWT(tokenString)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"token": newToken})
}

// Logout godoc
// @Summary Logout user
// @Tags Auth
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Router /auth/logout [post]
func (h *AuthHelper) Logout(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
