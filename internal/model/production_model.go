package model

import (
	"time"
)

type ProductionResponse struct {
	ID        uint            `json:"id"`
	UserID    uint            `json:"user_id,omitempty"`
	User      *UserResponse   `json:"user,omitempty"`
	OreID     uint16          `json:"ore_id,omitempty"`
	Ore       *OreResponse    `json:"ore,omitempty"`
	SourceID  uint            `json:"source_id,omitempty"`
	Source    *SourceResponse `json:"source,omitempty"`
	Weight    float32         `json:"weight"`
	Notes     *string         `json:"notes"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type CreateProductionRequest struct {
	SourceID uint    `json:"source_id" validate:"required,exists=sources.id"`
	OreID    uint16  `json:"ore_id"  validate:"required,exists=ores.id"`
	Weight   float32 `json:"weight" validate:"required,min=0.01,max=999999.99"`
	Notes    *string `json:"notes"`
}

type UpdateProductionRequest struct {
	SourceID uint    `json:"source_id" validate:"required,exists=sources.id"`
	OreID    uint16  `json:"ore_id"  validate:"required,exists=ores.id"`
	Weight   float32 `json:"weight" validate:"required,min=0.01,max=999999.99"`
	Notes    *string `json:"notes"`
}
