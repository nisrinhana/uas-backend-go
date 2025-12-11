package route

import (
    "github.com/gin-gonic/gin"
    "uas-backend-go/helper"
)

func LecturerRoutes(r *gin.Engine, h *helper.LecturerHelper) {
    g := r.Group("/api/v1/lecturers")

    g.GET("", h.GetAll)
    g.GET("/:id/advisees", h.GetAdvisees)
}
