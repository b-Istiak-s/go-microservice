package repository

import (
	"myapp/internal/model"

	"gorm.io/gorm"
)

// UserRepository defines the repository for user-related DB operations.
type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

// userRepository is the implementation of UserRepository.
type userRepository struct {
	db *gorm.DB
}

// Constructor for userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// Create inserts a new user into the database.
func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// FindByEmail finds a user by email.
func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
