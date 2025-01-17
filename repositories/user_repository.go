package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"datingApp/models"
)

type UserRepository interface {
	GetUserByID(userID uuid.UUID) (*models.User, error)
	GetUnswipedUsers(userID uuid.UUID, swipedIDs []uuid.UUID) ([]models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error // New method to create user
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, userID).Error
	return &user, err
}

func (r *UserRepo) GetUnswipedUsers(userID uuid.UUID, swipedIDs []uuid.UUID) ([]models.User, error) {
	var users []models.User

	query := r.DB.Where("id != ?", userID)
	if len(swipedIDs) > 0 {
		query = query.Where("id NOT IN ?", swipedIDs)
	}

	err := query.Find(&users).Error
	return users, err
}

func (r *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// New method to create a new user
func (r *UserRepo) CreateUser(user *models.User) error {
	result := r.DB.Create(user)
	return result.Error
}
