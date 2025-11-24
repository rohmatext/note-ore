package model

import "time"

type SourceResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber *string   `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateSourceRequest struct {
	Name        string  `json:"name" validate:"required"`
	PhoneNumber *string `json:"phone_number"`
}

type UpdateSourceRequest struct {
	Name        string  `json:"name" validate:"required"`
	PhoneNumber *string `json:"phone_number"`
}
