package models

import "github.com/google/uuid"

type SignUpResponse struct {
	UserID uuid.UUID `json:"user_id"`
}

type LoginResponse struct {
	UserID uuid.UUID `json:"user_id"`
	Token  string    `json:"token"`
}
