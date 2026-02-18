package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/auth"
	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/config"
	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/middlewares"
)

type Dependencies struct {
	Logger      zerolog.Logger
	AuthHandler *auth.Handler
	JWTManager  *auth.JWTManager
	Config      *config.Config
}

func NewRouter(dep Dependencies) *gin.Engine {
	if dep.Config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	if dep.Config.TrustedProxies() != nil {
		router.SetTrustedProxies(dep.Config.TrustedProxies())
	} else {
		router.SetTrustedProxies(nil)
	}

	router.Use(gin.Recovery())
	router.Use(middlewares.LoggerMiddleware(dep.Logger))
	router.Use(middlewares.ErrorMiddleware())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	auth.RegisterRoutes(router, dep.AuthHandler)

	protected := router.Group("/api")
	protected.Use(auth.AuthMiddleware(dep.JWTManager))

	return router
}
