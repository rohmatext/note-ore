package entity

import (
	"time"

	"gorm.io/gorm"
)

type Production struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	UserID    uint           `gorm:"column:user_id"`
	SourceID  uint           `gorm:"column:source_id"`
	OreID     uint16         `gorm:"column:ore_id"`
	Weight    float32        `gorm:"column:weight"`
	Notes     *string        `gorm:"column:notes"`
	User      *User          `gorm:"foreignKey:UserID;references:ID"`
	Source    *Source        `gorm:"foreignKey:SourceID;references:ID"`
	Ore       *Ore           `gorm:"foreignKey:OreID;references:ID"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Production) TableName() string {
	return "productions"
}
