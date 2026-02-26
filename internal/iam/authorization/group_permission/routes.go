package group_permission

import "github.com/gin-gonic/gin"

func RegisterRoutes(public, private *gin.RouterGroup, handler *Handler) {
	actions := private.Group("/actions")
	actions.GET("/", handler.List)
}
