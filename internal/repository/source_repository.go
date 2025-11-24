package repository

import (
	"rohmatext/ore-note/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SourceRepository interface {
	FindAll(db *gorm.DB) ([]*entity.Source, error)
	FindById(db *gorm.DB, id uint) (*entity.Source, error)
	Create(db *gorm.DB, source *entity.Source) error
	Update(db *gorm.DB, source *entity.Source) error
	Delete(db *gorm.DB, id uint) error
}

type SourceRepositoryImpl struct {
	Log *logrus.Logger
}

func NewSourceRepository(log *logrus.Logger) SourceRepository {
	return &SourceRepositoryImpl{
		Log: log,
	}
}

func (r *SourceRepositoryImpl) FindAll(db *gorm.DB) ([]*entity.Source, error) {
	var sources []*entity.Source
	err := db.Order("name ASC").Find(&sources).Error
	return sources, err
}

func (r *SourceRepositoryImpl) FindById(db *gorm.DB, id uint) (*entity.Source, error) {
	var source entity.Source
	if err := db.Where("id = ?", id).First(&source).Error; err != nil {
		return nil, err
	}
	return &source, nil

}

func (r *SourceRepositoryImpl) Create(db *gorm.DB, source *entity.Source) error {
	return db.Create(&source).Error
}

func (r *SourceRepositoryImpl) Update(db *gorm.DB, source *entity.Source) error {
	return db.Save(&source).Error
}

func (r *SourceRepositoryImpl) Delete(db *gorm.DB, id uint) error {
	var source entity.Source
	return db.Where("id = ?", id).Delete(&source).Error
}
