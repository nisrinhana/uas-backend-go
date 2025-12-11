package helper

import (
    "github.com/gin-gonic/gin"
    "uas-backend-go/app/service"
)

type StudentHelper struct {
    Service *service.StudentService
}

func NewStudentHelper(s *service.StudentService) *StudentHelper {
    return &StudentHelper{Service: s}
}

func (h *StudentHelper) GetAll(c *gin.Context) {
    data, err := h.Service.GetAll(c.Request.Context())
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, data)
}

func (h *StudentHelper) GetByID(c *gin.Context) {
    id := c.Param("id")
    data, err := h.Service.GetByID(c.Request.Context(), id)
    if err != nil {
        c.JSON(404, gin.H{"error": "student not found"})
        return
    }
    c.JSON(200, data)
}

func (h *StudentHelper) GetAchievements(c *gin.Context) {
    id := c.Param("id")
    data, err := h.Service.GetAchievements(c.Request.Context(), id)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, data)
}

func (h *StudentHelper) UpdateAdvisor(c *gin.Context) {
    id := c.Param("id")

    var body struct {
        AdvisorID *string `json:"advisor_id"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(400, gin.H{"error": "invalid body"})
        return
    }

    err := h.Service.UpdateAdvisor(c.Request.Context(), id, body.AdvisorID)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"message": "advisor updated"})
}
