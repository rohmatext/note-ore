package model

import "time"

type OreResponse struct {
	ID        uint16    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateOreRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateOreRequest struct {
	Name string `json:"name" validate:"required"`
}
