package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"datingApp/models"
	"datingApp/repositories"
)

type PremiumServiceInterface interface {
	RegisterPremium(userID uuid.UUID, packageID uuid.UUID) error
	IsUserPremium(userID uuid.UUID) (bool, error)
	GetUserPremiumDetails(userID uuid.UUID) (*models.UserPremium, error)
	GetAllPremiumPackages() ([]models.PremiumPackage, error)
	GetPremiumPackageByID(packageID uuid.UUID) (*models.PremiumPackage, error)
	CreatePremiumPackage(name, description string, price decimal.Decimal) error
	UpdatePremiumPackage(pkg *models.PremiumPackage) error
	DeletePremiumPackage(packageID uuid.UUID) error
}

type PremiumService struct {
	premiumRepo repositories.PremiumRepository
}

func NewPremiumService(repo repositories.PremiumRepository) *PremiumService {
	return &PremiumService{
		premiumRepo: repo,
	}
}

// GetPremiumPackageByID retrieves a specific premium package by ID
func (s *PremiumService) GetPremiumPackageByID(packageID uuid.UUID) (*models.PremiumPackage, error) {
	return s.premiumRepo.GetPremiumPackageByID(packageID)
}

// RegisterPremium registers a new premium subscription for a user
func (s *PremiumService) RegisterPremium(userID uuid.UUID, packageID uuid.UUID) error {
	// Validate package exists
	pkg, err := s.GetPremiumPackageByID(packageID)
	if err != nil {
		return err
	}
	if pkg == nil {
		return errors.New("invalid premium package")
	}

	// Check if user already has an active premium subscription
	isPremium, err := s.IsUserPremium(userID)
	if err != nil {
		return err
	}
	if isPremium {
		return errors.New("user already has an active premium subscription")
	}

	// Create new premium subscription
	userPremium := &models.UserPremium{
		UserID:       userID,
		PackageID:    packageID,
		PurchaseDate: time.Now(),
	}

	return s.premiumRepo.RegisterPremium(userPremium)
}

// IsUserPremium checks if a user has an active premium subscription
func (s *PremiumService) IsUserPremium(userID uuid.UUID) (bool, error) {
	return s.premiumRepo.IsUserPremium(userID)
}

// GetUserPremiumDetails retrieves detailed information about a user's premium subscription
func (s *PremiumService) GetUserPremiumDetails(userID uuid.UUID) (*models.UserPremium, error) {
	return s.premiumRepo.GetUserPremium(userID)
}

// GetAllPremiumPackages retrieves all available premium packages
func (s *PremiumService) GetAllPremiumPackages() ([]models.PremiumPackage, error) {
	return s.premiumRepo.GetPremiumPackages()
}

// CreatePremiumPackage creates a new premium package
func (s *PremiumService) CreatePremiumPackage(name, description string, price decimal.Decimal) error {
	if price.IsNegative() {
		return errors.New("price cannot be negative")
	}

	pk := &models.PremiumPackage{
		PackageName: name,
		Description: description,
		Price:       price,
	}

	return s.premiumRepo.CreatePremiumPackage(pk)
}

// UpdatePremiumPackage updates an existing premium package
func (s *PremiumService) UpdatePremiumPackage(pkg *models.PremiumPackage) error {
	if pkg.Price.IsNegative() {
		return errors.New("price cannot be negative")
	}

	existing, err := s.GetPremiumPackageByID(pkg.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("premium package not found")
	}

	return s.premiumRepo.UpdatePremiumPackage(pkg)
}

// DeletePremiumPackage deletes a premium package
func (s *PremiumService) DeletePremiumPackage(packageID uuid.UUID) error {
	existing, err := s.GetPremiumPackageByID(packageID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("premium package not found")
	}

	return s.premiumRepo.DeletePremiumPackage(packageID)
}
