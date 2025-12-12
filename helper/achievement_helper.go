package helper

import (
    "net/http"
    "time"

    "uas-backend-go/app/model"
    "uas-backend-go/app/service"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type AchievementHelper struct {
    Service *service.AchievementService
}

func NewAchievementHelper(s *service.AchievementService) *AchievementHelper {
    return &AchievementHelper{Service: s}
}

// GetAllAchievements godoc
// @Summary Get achievements (role-based)
// @Tags Achievements
// @Security BearerAuth
// @Router /achievements [get]
func (h *AchievementHelper) GetAll(c *gin.Context) {
    role := c.GetString("role")
    userID := c.GetString("user_id")
    ctx := c.Request.Context()

    switch role {

    // ================================
    // STUDENT: return ref + mongo data
    // ================================
    case "student":
        var studentID string
        err := h.Service.RefRepo.DB.QueryRow(ctx,
            "SELECT id FROM students WHERE user_id = $1", userID).Scan(&studentID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "student not found"})
            return
        }

        // Ambil reference dari PostgreSQL
        refs, err := h.Service.GetForStudent(ctx, studentID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        var result []gin.H

        for _, ref := range refs {
            objID, err := primitive.ObjectIDFromHex(ref.MongoAchievementID)
            if err != nil {
                continue
            }

            ach, err := h.Service.GetByID(ctx, objID)
            if err != nil {
                continue
            }

            result = append(result, gin.H{
                "reference":   ref,
                "achievement": ach,
            })
        }

        c.JSON(http.StatusOK, result)
        return

    // ================================
    // LECTURER
    // ================================
    case "lecturer":
        refs, err := h.Service.GetForAdvisor(ctx, userID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, refs)
        return

    // ================================
    // ADMIN
    // ================================
    case "admin":
        refs, err := h.Service.GetAll(ctx)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, refs)
        return

    default:
        c.JSON(http.StatusForbidden, gin.H{"error": "role not allowed"})
        return
    }
}


// GetAchievementDetail godoc
// @Summary Get achievement detail
// @Tags Achievements
// @Security BearerAuth
// @Param id path string true "Mongo Achievement ID"
// @Router /achievements/{id} [get]
func (h *AchievementHelper) GetDetail(c *gin.Context) {
    mongoID := c.Param("id")

    objID, err := primitive.ObjectIDFromHex(mongoID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    ctx := c.Request.Context()
    ach, err := h.Service.GetByID(ctx, objID)

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "achievement not found"})
        return
    }

    c.JSON(http.StatusOK, ach)
}

// CreateAchievement godoc
// @Summary Create new achievement draft
// @Tags Achievements
// @Security BearerAuth
// @Accept json
// @Router /achievements [post]
func (h *AchievementHelper) Create(c *gin.Context) {
    var req model.Achievement

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    ctx := c.Request.Context()
    userID := c.GetString("user_id")

    // Ambil student ID dari tabel students berdasarkan user_id
    var studentID string
    err := h.Service.RefRepo.DB.QueryRow(ctx,
        "SELECT id FROM students WHERE user_id=$1", userID).Scan(&studentID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "student not found"})
        return
    }

    refID, err := h.Service.Create(ctx, studentID, req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

   c.JSON(http.StatusCreated, gin.H{
    "message": "achievement draft created",
    "reference_id": refID,
    "achievement": req, 
})

}

// UpdateAchievement godoc
// @Summary Update draft achievement
// @Tags Achievements
// @Security BearerAuth
// @Param id path string true "Mongo Achievement ID"
// @Accept json
// @Router /achievements/{id} [put]
func (h *AchievementHelper) Update(c *gin.Context) {
    mongoID := c.Param("id")

    var req model.Achievement
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    ctx := c.Request.Context()
    err := h.Service.Update(ctx, mongoID, req)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

  objID, _ := primitive.ObjectIDFromHex(mongoID)
    updated, err := h.Service.GetByID(ctx, objID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch updated data"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":     "achievement updated",
        "achievement": updated,
    })
}


// DeleteAchievement godoc
// @Summary Delete achievement draft
// @Tags Achievements
// @Security BearerAuth
// @Param id path string true "Mongo Achievement ID"
// @Router /achievements/{id} [delete]
func (h *AchievementHelper) Delete(c *gin.Context) {
    mongoID := c.Param("id")
    ctx := c.Request.Context()
    userID := c.GetString("user_id")

    // 1. Ambil student_id
    var studentID string
    err := h.Service.RefRepo.DB.QueryRow(ctx,
        "SELECT id FROM students WHERE user_id = $1", userID).Scan(&studentID)
    if err != nil {
        c.JSON(http.StatusForbidden, gin.H{"error": "student not found"})
        return
    }

    // 2. Cari reference berdasarkan mongo id
    ref, err := h.Service.GetRefByMongoID(ctx, mongoID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "achievement reference not found"})
        return
    }

    // 3. Validasi kepemilikan
    if ref.StudentID != studentID {
        c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
        return
    }

    // 4. Validasi status draft
    if ref.Status != "draft" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "only draft can be deleted"})
        return
    }

    // 5. Hapus + dapatkan snapshot
    snapshot, err := h.Service.Delete(ctx, ref.ID, mongoID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 6. Return lengkap
    c.JSON(http.StatusOK, gin.H{
        "message":  "achievement deleted",
        "ref_id":   ref.ID,
        "mongo_id": mongoID,
        "status":   "deleted",
        "snapshot": snapshot,
    })
}


// SubmitAchievement godoc
// @Summary Submit achievement for verification
// @Tags Achievements
// @Security BearerAuth
// @Param id path string true "Reference ID"
// @Router /achievements/{id}/submit [post]
func (h *AchievementHelper) Submit(c *gin.Context) {
    refID := c.Param("id")
    ctx := c.Request.Context()

    // 1. Ambil user_id dari token
    userID := c.GetString("user_id")

    // 2. Ambil student_id berdasarkan user_id
    var studentID string
    err := h.Service.RefRepo.DB.QueryRow(ctx,
        "SELECT id FROM students WHERE user_id = $1", userID).Scan(&studentID)
    if err != nil {
        c.JSON(http.StatusForbidden, gin.H{"error": "student not found"})
        return
    }

    // 3. Ambil reference dari DB
    ref, err := h.Service.GetRefByID(ctx, refID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }

    // 4. Validasi kepemilikan
    if ref.StudentID != studentID {
        c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
        return
    }

    // 5. Validasi status
    if ref.Status != "draft" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "must be draft"})
        return
    }

    // 6. Update status → submitted
    err = h.Service.Submit(ctx, ref)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "submitted"})
}


// VerifyAchievement godoc
// @Summary Verify student achievement (Lecturer)
// @Tags Achievements
// @Security BearerAuth
// @Param id path string true "Reference ID"
// @Router /achievements/{id}/verify [post]
func (h *AchievementHelper) Verify(c *gin.Context) {
    refID := c.Param("id")
    userID := c.GetString("user_id")
    ctx := c.Request.Context()

    // ambil lecturer.id berdasarkan user_id
    var lecturerID string
    err := h.Service.RefRepo.DB.QueryRow(ctx,
        "SELECT id FROM lecturers WHERE user_id = $1", userID).Scan(&lecturerID)
    if err != nil {
        c.JSON(http.StatusForbidden, gin.H{"error": "lecturer not found"})
        return
    }

    ref, err := h.Service.GetRefByID(ctx, refID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }

    if ref.Status != "submitted" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "must be submitted"})
        return
    }

    ref.Status = "verified"

    err = h.Service.Verify(ctx, ref, lecturerID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "verified"})
}


// RejectAchievement godoc
// @Summary Reject student achievement
// @Tags Achievements
// @Security BearerAuth
// @Param id path string true "Reference ID"
// @Router /achievements/{id}/reject [post]
func (h *AchievementHelper) Reject(c *gin.Context) {
    refID := c.Param("id")
    ctx := c.Request.Context()

    var body struct {
        Note string `json:"note"`
    }
    c.ShouldBindJSON(&body)

    ref, err := h.Service.GetRefByID(ctx, refID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }

    if ref.Status != "submitted" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "must be submitted"})
        return
    }

    ref.Status = "rejected"

    err = h.Service.Reject(ctx, ref, body.Note)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "rejected"})
}


// GetAchievementHistory godoc
// @Summary Get achievement history
// @Tags Achievements
// @Security BearerAuth
// @Param id path string true "Reference ID"
// @Router /achievements/{id}/history [get]
func (h *AchievementHelper) GetHistory(c *gin.Context) {
    refID := c.Param("id")
    ctx := c.Request.Context()

    ref, err := h.Service.GetRefByID(ctx, refID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }

    history := gin.H{
        "status":        ref.Status,
        "submitted_at":  ref.SubmittedAt,
        "verified_at":   ref.VerifiedAt,
        "verified_by":   ref.VerifiedBy,
        "rejection_note": ref.RejectionNote,
    }

    c.JSON(http.StatusOK, history)
}


// UploadAttachment godoc
// @Summary Upload achievement attachment
// @Tags Achievements
// @Security BearerAuth
// @Param id path string true "Reference ID"
// @Param file formData file true "Upload file"
// @Router /achievements/{id}/attachments [post]
func (h *AchievementHelper) UploadAttachment(c *gin.Context) {
    refID := c.Param("id")              // <-- Ini harus refID (bukan mongoID)
    userID := c.GetString("user_id")
    ctx := c.Request.Context()

    // 1. ambil student_id berdasarkan user_id
    var studentID string
    err := h.Service.RefRepo.DB.QueryRow(ctx,
        "SELECT id FROM students WHERE user_id = $1", userID).Scan(&studentID)

    if err != nil {
        c.JSON(http.StatusForbidden, gin.H{"error": "student not found"})
        return
    }

    // 2. ambil reference berdasarkan refID
    ref, err := h.Service.GetRefByID(ctx, refID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "achievement reference not found"})
        return
    }

    // 3. validasi kepemilikan achievement
    if ref.StudentID != studentID {
        c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
        return
    }

    // 4. ambil file
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "file missing"})
        return
    }

    filename := "uploads/" + time.Now().Format("20060102150405") + "_" + file.Filename
    c.SaveUploadedFile(file, filename)

    c.JSON(http.StatusOK, gin.H{
        "message": "uploaded",
        "url": filename,
    })
}

