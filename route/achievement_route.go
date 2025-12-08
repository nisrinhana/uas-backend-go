package route

import (
    "github.com/gin-gonic/gin"
    "uas-backend-go/helper"
    "uas-backend-go/middleware"
)

func AchievementRoutes(r *gin.Engine, h *helper.AchievementHelper) {
    g := r.Group("/api/v1/achievements")

    g.GET("", middleware.AuthRequired("view_achievements"), h.GetAll)
    g.GET("/:id", middleware.AuthRequired("view_achievements"), h.GetDetail)

    g.POST("", middleware.AuthRequired("create_achievement"), h.Create)
    g.PUT("/:id", middleware.AuthRequired("edit_achievement"), h.Update)
    g.DELETE("/:id", middleware.AuthRequired("delete_achievement"), h.Delete)

    g.POST("/:id/submit", middleware.AuthRequired("submit_achievement"), h.Submit)
    g.POST("/:id/verify", middleware.AuthRequired("verify_achievement"), h.Verify)
    g.POST("/:id/reject", middleware.AuthRequired("verify_achievement"), h.Reject)

    g.GET("/:id/history", middleware.AuthRequired("view_achievements"), h.GetHistory)

    g.POST("/:id/attachments", middleware.AuthRequired("upload_attachment"), h.UploadAttachment)
}
