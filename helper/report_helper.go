package helper

import (
    "github.com/gin-gonic/gin"
    "uas-backend-go/app/service"
)

type ReportHelper struct {
    Service *service.ReportService
}

func NewReportHelper(s *service.ReportService) *ReportHelper {
    return &ReportHelper{Service: s}
}

// GetGlobalStatistics godoc
// @Summary Get global achievement statistics
// @Tags Reports
// @Security BearerAuth
// @Router /reports/statistics [get]
func (h *ReportHelper) GetGlobal(c *gin.Context) {
    data, err := h.Service.GetGlobalStatistics(c.Request.Context())
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, data)
}

// GetStudentStatistics godoc
// @Summary Get statistics for a student
// @Tags Reports
// @Security BearerAuth
// @Param id path string true "Student ID"
// @Router /reports/student/{id} [get]
func (h *ReportHelper) GetStudent(c *gin.Context) {
    id := c.Param("id")

    data, err := h.Service.GetStudentStatistics(c.Request.Context(), id)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, data)
}
