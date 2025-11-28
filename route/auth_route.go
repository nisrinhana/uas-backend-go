package route

import (
    "github.com/gin-gonic/gin"
    "uas-backend-go/helper"
    "uas-backend-go/middleware"
)

func AuthRoutes(r *gin.Engine, h *helper.AuthHelper) {
    g := r.Group("/api/v1/auth")
    {
        g.POST("/login", h.Login)
        g.GET("/profile", middleware.AuthRequired("view_profile"), h.Profile)
    }
}
