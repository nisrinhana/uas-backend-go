package middleware

import (
    "context"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "uas-backend-go/database"
    "uas-backend-go/utils"
)

func AuthRequired(permission string) gin.HandlerFunc {
    return func(c *gin.Context) {

        auth := c.GetHeader("Authorization")
        if !strings.HasPrefix(auth, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(auth, "Bearer ")

        claims, err := utils.VerifyJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }

        // --- Check Permission ---
        has := false
        for _, p := range claims.Permissions {
            if p == permission {
                has = true
                break
            }
        }
        if !has {
            c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
            c.Abort()
            return
        }

        // --- NEW: ambil nama role dari DB ---
        var roleName string
        err = database.Postgres.QueryRow(context.Background(),
            "SELECT name FROM roles WHERE id = $1", claims.RoleID).Scan(&roleName)
        if err != nil {
            c.JSON(http.StatusForbidden, gin.H{"error": "role not allowed"})
            c.Abort()
            return
        }

        // simpan ke context
        c.Set("user_id", claims.UserID)
        c.Set("role_id", claims.RoleID)
        c.Set("permissions", claims.Permissions)
        c.Set("role", roleName)

        c.Next()
    }
}
