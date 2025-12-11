package route

import (
    "uas-backend-go/helper"
    "github.com/gin-gonic/gin"
)

func StudentRoutes(r *gin.Engine, h *helper.StudentHelper) {
    g := r.Group("/api/v1/students")

    g.GET("", h.GetAll)
    g.GET("/:id", h.GetByID)
    g.GET("/:id/achievements", h.GetAchievements)
    g.PUT("/:id/advisor", h.UpdateAdvisor)
}
