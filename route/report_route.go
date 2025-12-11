package route

import (
    "github.com/gin-gonic/gin"
    "uas-backend-go/helper"
)

func ReportRoutes(r *gin.Engine, h *helper.ReportHelper) {
    g := r.Group("/api/v1/reports")

    g.GET("/statistics", h.GetGlobal)
    g.GET("/student/:id", h.GetStudent)
}
