package organizations

import "github.com/gin-gonic/gin"

func RegisterRoutes(public, private *gin.RouterGroup, handler *Handler) {
	organizations := private.Group("/organizations")
	organizations.GET("/", handler.List)
	organizations.GET("/:id", handler.Details)
	organizations.POST("/", handler.Create)
	organizations.PUT("/:id", handler.Update)
	organizations.DELETE("/:id", handler.Delete)
}
