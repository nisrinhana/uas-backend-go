package middleware

import (
    "net/http"
    "strings"
    "uas-backend-go/helper"

    "github.com/gin-gonic/gin"
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

        claims, err := helper.VerifyJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }

        // Check permission
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

        // inject ke gin
        c.Set("user_id", claims.UserID)
        c.Set("role_id", claims.RoleID)
        c.Set("permissions", claims.Permissions)

        c.Next()
    }
}
