package route

import (
	controller "myapp/internal/controller/auth"
	"myapp/internal/db"
	"myapp/internal/middleware"
	repository "myapp/internal/repository/auth"
	service "myapp/internal/service/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	// Set up repository
	userRepo := repository.NewUserRepository(db.DB)

	// Set up service
	authService := service.NewAuthService(userRepo)

	// Set up controller
	authController := controller.NewAuthController(authService)

	// Group your routes under /api/auth
	// authRoutes := router.Group("/api/auth") // for development
	authRoutes := router.Group("/") // for k8s
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/verify", middleware.AuthMiddleware(), authController.UserExists)
	}
}
