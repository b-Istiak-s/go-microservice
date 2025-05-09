package controller

import (
	"fmt"
	service "myapp/internal/service/auth"
	"myapp/internal/util/response"
	validator "myapp/internal/validator/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController handles authentication routes.
type AuthController struct {
	authService service.AuthService
}

// NewAuthController creates a new AuthController.
func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService}
}

// Register handles user registration.
func (authController *AuthController) Register(context *gin.Context) {
	var req validator.RegisterRequest

	// Bind and validate request
	ok, errs := validator.BindAndValidateRegistration(context, &req)
	if !ok {
		response.Error(context, http.StatusUnprocessableEntity, "Validation Error", errs)
		return
	}

	// Call service
	user, err := authController.authService.Register(req)
	if err != nil {
		response.Error(context, http.StatusBadRequest, fmt.Sprintf("Error: %v", err.Error()))
		return
	}

	response.Success(context, http.StatusCreated, "Registration Successful", user)
}

// Login handles user login.
func (authController *AuthController) Login(context *gin.Context) {
	var req validator.LoginRequest

	ok, errs := validator.BindAndValidateLogin(context, &req)
	if !ok {
		response.Error(context, http.StatusUnprocessableEntity, "Validation Error", errs)
		return
	}

	// Call service
	token, err := authController.authService.Login(req)
	if err != nil {
		response.Error(context, http.StatusBadRequest, fmt.Sprintf("Error: %v", err.Error()))
		return
	}

	response.LoginSuccess(context, http.StatusOK, "Login Successful", token)
}

// Checks if the user exists
func (authController *AuthController) UserExists(context *gin.Context) {

	userIDInterface, exists := context.Get("userID")
	if !exists {
		response.Error(context, http.StatusUnauthorized, "User ID not found in context")
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		response.Error(context, http.StatusInternalServerError, "Invalid user ID type")
	}

	_, err := authController.authService.UserExists(userID)
	if err != nil {
		response.Error(context, http.StatusBadRequest, fmt.Sprintf("Error: %v", err.Error()))
		return
	}

	response.Success(context, http.StatusOK, "User Exists Check")
}
