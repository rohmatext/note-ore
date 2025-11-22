package repository

import (
	"rohmatext/ore-note/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OreRepository interface {
	FindAll(db *gorm.DB) ([]*entity.Ore, error)
	FindById(db *gorm.DB, id uint16) (*entity.Ore, error)
	Create(db *gorm.DB, ore *entity.Ore) error
	Update(db *gorm.DB, ore *entity.Ore) error
	Delete(db *gorm.DB, id uint16) error
}

type OreRepositoryImpl struct {
	Log *logrus.Logger
}

func NewOreRepository(log *logrus.Logger) OreRepository {
	return &OreRepositoryImpl{
		Log: log,
	}
}

func (r *OreRepositoryImpl) FindAll(db *gorm.DB) ([]*entity.Ore, error) {
	var ores []*entity.Ore
	err := db.Order("name ASC").Find(&ores).Error
	return ores, err
}

func (r *OreRepositoryImpl) FindById(db *gorm.DB, id uint16) (*entity.Ore, error) {
	var ore entity.Ore
	if err := db.Where("id = ?", id).First(&ore).Error; err != nil {
		return nil, err
	}
	return &ore, nil

}

func (r *OreRepositoryImpl) Create(db *gorm.DB, ore *entity.Ore) error {
	return db.Create(&ore).Error
}

func (r *OreRepositoryImpl) Update(db *gorm.DB, ore *entity.Ore) error {
	return db.Save(&ore).Error
}

func (r *OreRepositoryImpl) Delete(db *gorm.DB, id uint16) error {
	var ore entity.Ore
	return db.Where("id = ?", id).Delete(&ore).Error
}
