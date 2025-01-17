package models

import "github.com/google/uuid"

type SignUpRequest struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	Username      string `json:"username"`
	ProfilePicURL string `json:"profilePicURL"`
	Bio           string `json:"bio"`
	Interests     string `json:"interests"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PremiumPackageRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       string `json:"price" binding:"required"`
}

type PurchaseRequest struct {
	UserID    string `json:"user_id" binding:"required"`
	PackageID string `json:"package_id" binding:"required"`
}

type SwipeRequest struct {
	ProfileID uuid.UUID `json:"profile_id" binding:"required"`
}
