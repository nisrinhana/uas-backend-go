package route

import (
    "github.com/gin-gonic/gin"
    "uas-backend-go/helper"
    "uas-backend-go/middleware"
)

func UserRoutes(r *gin.Engine, h *helper.UserHelper) {
    g := r.Group("/api/v1/users")
    {
        g.GET("", middleware.AuthRequired("view_users"), h.GetAll)
        g.GET("/:id", middleware.AuthRequired("view_users"), h.GetByID)

        g.POST("", middleware.AuthRequired("manage_users"), h.Create)

        g.PUT("/:id", middleware.AuthRequired("edit_users"), h.Update)
        g.DELETE("/:id", middleware.AuthRequired("delete_users"), h.Delete)
        g.PUT("/:id/role", middleware.AuthRequired("edit_users"), h.UpdateRole)
    }
}
