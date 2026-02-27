package group_permission

import "github.com/gin-gonic/gin"

func RegisterRoutes(public, private *gin.RouterGroup, handler *Handler) {
	groupPermissions := private.Group("/group-permissions")
	groupPermissions.GET("/", handler.List)
	groupPermissions.GET("/:id", handler.Details)
	groupPermissions.POST("/", handler.Create)
	groupPermissions.PUT("/:id", handler.Update)
	groupPermissions.DELETE("/:id", handler.Delete)
}
