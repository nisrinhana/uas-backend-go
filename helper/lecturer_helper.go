package helper

import (
    "github.com/gin-gonic/gin"
    "uas-backend-go/app/service"
)

type LecturerHelper struct {
    Service *service.LecturerService
}

func NewLecturerHelper(s *service.LecturerService) *LecturerHelper {
    return &LecturerHelper{Service: s}
}

// GetAllLecturers godoc
// @Summary Get all lecturers
// @Tags Lecturers
// @Security BearerAuth
// @Success 200 {array} model.Lecturer
// @Router /lecturers [get]
func (h *LecturerHelper) GetAll(c *gin.Context) {
    data, err := h.Service.GetAll(c.Request.Context())
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, data)
}


// GetAdvisees godoc
// @Summary Get lecturer advisees
// @Tags Lecturers
// @Security BearerAuth
// @Param id path string true "Lecturer ID"
// @Router /lecturers/{id}/advisees [get]
func (h *LecturerHelper) GetAdvisees(c *gin.Context) {
    id := c.Param("id")
    data, err := h.Service.GetAdvisees(c.Request.Context(), id)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, data)
}
