package helper

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "uas-backend-go/app/model"
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


// POST /api/v1/auth/refresh
func (h *AuthHelper) Refresh(c *gin.Context) {
    authHeader := c.GetHeader("Authorization")
    if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
        return
    }

    tokenString := authHeader[7:]
    newToken, err := RefreshJWT(tokenString) // langsung panggil fungsi
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": newToken})
}


// POST /api/v1/auth/logout
func (h *AuthHelper) Logout(c *gin.Context) {
    // stateless JWT logout, hanya kasih pesan
    c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
