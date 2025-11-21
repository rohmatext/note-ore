package model

import (
	"time"
)

type RefreshTokenResponse struct {
	ID        uint      `json:"id,omitempty"`
	Token     string    `json:"token,omitempty"`
	UserId    uint      `json:"user_id,omitempty"`
	ExpiresAt time.Time `json:"expired_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
