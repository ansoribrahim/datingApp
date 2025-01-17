package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"datingApp/models"
)

type PremiumRepository interface {
	RegisterPremium(userPremium *models.UserPremium) error
	IsUserPremium(userID uuid.UUID) (bool, error)
	GetUserPremium(userID uuid.UUID) (*models.UserPremium, error)
	GetPremiumPackages() ([]models.PremiumPackage, error)
	GetPremiumPackageByID(packageID uuid.UUID) (*models.PremiumPackage, error)
	CreatePremiumPackage(pkg *models.PremiumPackage) error
	UpdatePremiumPackage(pkg *models.PremiumPackage) error
	DeletePremiumPackage(packageID uuid.UUID) error
}

type PremiumRepo struct {
	DB *gorm.DB
}

func NewPremiumRepo(db *gorm.DB) *PremiumRepo {
	return &PremiumRepo{DB: db}
}

// RegisterPremium registers a new premium subscription for a user
func (r *PremiumRepo) RegisterPremium(userPremium *models.UserPremium) error {
	return r.DB.Create(userPremium).Error
}

// IsUserPremium checks if a user has an active premium subscription
func (r *PremiumRepo) IsUserPremium(userID uuid.UUID) (bool, error) {
	var userPremium models.UserPremium
	err := r.DB.Where("user_id = ?", userID).
		Order("purchase_date DESC").
		First(&userPremium).Error

	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	// Check if the subscription is still valid (you might want to add duration logic here)
	return true, nil
}

// GetUserPremium retrieves the user's premium subscription details
func (r *PremiumRepo) GetUserPremium(userID uuid.UUID) (*models.UserPremium, error) {
	var userPremium models.UserPremium
	err := r.DB.Where("user_id = ?", userID).
		Order("purchase_date DESC").
		First(&userPremium).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &userPremium, nil
}

// GetPremiumPackages retrieves all available premium packages
func (r *PremiumRepo) GetPremiumPackages() ([]models.PremiumPackage, error) {
	var packages []models.PremiumPackage
	err := r.DB.Find(&packages).Error
	return packages, err
}

// GetPremiumPackageByID retrieves a specific premium package by ID
func (r *PremiumRepo) GetPremiumPackageByID(packageID uuid.UUID) (*models.PremiumPackage, error) {
	var pkg models.PremiumPackage
	err := r.DB.Where("package_id = ?", packageID).First(&pkg).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &pkg, err
}

// CreatePremiumPackage creates a new premium package
func (r *PremiumRepo) CreatePremiumPackage(pkg *models.PremiumPackage) error {
	return r.DB.Create(pkg).Error
}

// UpdatePremiumPackage updates an existing premium package
func (r *PremiumRepo) UpdatePremiumPackage(pkg *models.PremiumPackage) error {
	return r.DB.Save(pkg).Error
}

// DeletePremiumPackage deletes a premium package
func (r *PremiumRepo) DeletePremiumPackage(packageID uuid.UUID) error {
	return r.DB.Where("package_id = ?", packageID).Delete(&models.PremiumPackage{}).Error
}
