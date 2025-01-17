package services

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"

	"datingApp/models"
	"datingApp/repositories"
)

type AuthService struct {
	UserRepo  repositories.UserRepository
	SecretKey string
}

func NewAuthService(userRepo repositories.UserRepository, secretKey string) *AuthService {
	return &AuthService{UserRepo: userRepo, SecretKey: secretKey}
}

func (s *AuthService) SignUp(req models.SignUpRequest) (*models.SignUpResponse, error) {
	existingUser, err := s.UserRepo.GetUserByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already in use")
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &models.User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
	}

	err = s.UserRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	// Return only the user ID in a SignUpResponse struct
	return &models.SignUpResponse{UserID: user.ID}, nil
}

func (s *AuthService) Login(email, password string) (*models.LoginResponse, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !checkPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	// Return user ID and token in a LoginResponse struct
	return &models.LoginResponse{
		UserID: user.ID,
		Token:  token,
	}, nil
}

func (s *AuthService) generateToken(user *models.User) (string, error) {
	// Create the JWT claims, which includes the username and expiration time
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Set token expiration (e.g., 72 hours)
	}

	// Create token with claims and sign it with your secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
