package route

import (
	controller "myapp/internal/controller/auth"
	"myapp/internal/db"
	repository "myapp/internal/repository/auth"
	service "myapp/internal/service/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Set up repository
	userRepo := repository.NewUserRepository(db.DB)

	// Set up service
	authService := service.NewAuthService(userRepo)

	// Set up controller
	authController := controller.NewAuthController(authService)

	// Group your routes under /api/auth
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", authController.Register)
		// authRoutes.POST("/login", authController.Login) // Enable when Login is ready
	}
}

func LoginRoutes(router *gin.Engine) {
	// Set up repository
	userRepo := repository.NewUserRepository(db.DB)

	// Set up service
	authService := service.NewAuthService(userRepo)

	// Set up controller
	authController := controller.NewAuthController(authService)

	// Group your routes under /api/auth
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/login", authController.Login)
	}
}
