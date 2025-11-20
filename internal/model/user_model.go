package model

import "time"

type UserResponse struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required"`
}
