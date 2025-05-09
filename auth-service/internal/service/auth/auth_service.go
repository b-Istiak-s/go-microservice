package service

import (
	"errors"
	"myapp/internal/auth"
	"myapp/internal/model"
	repository "myapp/internal/repository/auth"
	validator "myapp/internal/validator/auth"
)

// AuthService defines the service for authentication.
type AuthService interface {
	Register(req validator.RegisterRequest) (*model.User, error)
	Login(req validator.LoginRequest) (string, error)
	UserExists(userID uint) (bool, error)
}

// authService is the implementation of AuthService.
type authService struct {
	userRepository repository.UserRepository
}

// Constructor for authService
func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}

func (authS *authService) Register(req validator.RegisterRequest) (*model.User, error) {
	// Check if user already exists
	existingUser, _ := authS.userRepository.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("Email already registered")
	}

	// Hash the password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create new user
	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := authS.userRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (authS *authService) Login(req validator.LoginRequest) (string, error) {
	// Check if user exists
	user, err := authS.userRepository.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Check password
	if !auth.CheckPassword(user.Password, req.Password) {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func (authS *authService) UserExists(userID uint) (bool, error) {
	// Check if user exists
	user, err := authS.userRepository.FindByID(userID)
	if err != nil {
		return false, err
	}

	return user != nil, nil
}
