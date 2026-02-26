package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/config"
	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/iam/identity/auth"
	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/iam/identity/user"
	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/iam/provisioning/token"
	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/middlewares"
	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/notifications/emails"
)

type Dependencies struct {
	Logger zerolog.Logger
	Config *config.Config
	DB     *gorm.DB
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
	router.Use(middlewares.I18nMiddleware())
	router.Use(middlewares.LoggerMiddleware(dep.Logger))
	router.Use(middlewares.ErrorMiddleware())

	public := router.Group("")
	private := router.Group("")

	jwtManager := auth.NewJWTManager(dep.Config.AccessTokenSecret, dep.Config.RefreshTokenSecret)

	// Middlewares
	private.Use(auth.AuthMiddleware(jwtManager))

	// Routes
	handlers := setupHandlers(dep, jwtManager)
	setupRoutes(public, private, handlers)

	return router
}

type Handlers struct {
	AuthHandler *auth.Handler
	UserHandler *user.Handler
}

func setupHandlers(dep Dependencies, jwtManager *auth.JWTManager) *Handlers {
	emailSender := emails.NewSMTPEmailSender(
		dep.Config.SMTPHost,
		dep.Config.SMTPPort,
		dep.Config.SMTPUser,
		dep.Config.SMTPPassword,
		dep.Config.SMTPFrom,
	)

	userRepository := user.NewGormRepository(dep.DB)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	tokenRepository := token.NewGormRepository(dep.DB)
	tokenService := token.NewService(tokenRepository)

	authService := auth.NewService(userRepository, jwtManager, tokenService, emailSender, dep.Config.FrontendBaseURL)
	authHandler := auth.NewHandler(authService)

	return &Handlers{
		AuthHandler: authHandler,
		UserHandler: userHandler,
	}
}

func setupRoutes(public, private *gin.RouterGroup, handlers *Handlers) {
	auth.RegisterRoutes(public, private, handlers.AuthHandler)
	user.RegisterRoutes(public, private, handlers.UserHandler)

	public.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})
}
