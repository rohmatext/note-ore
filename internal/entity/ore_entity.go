package entity

import (
	"time"

	"gorm.io/gorm"
)

type Ore struct {
	ID        uint16         `gorm:"column:id;primaryKey"`
	Name      string         `gorm:"column:name"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Ore) TableName() string {
	return "ores"
}
