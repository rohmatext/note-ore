package repository

import (
	"rohmatext/ore-note/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductionRepository interface {
	FindAll(db *gorm.DB) ([]*entity.Production, error)
	FindByMonth(db *gorm.DB, year uint, month uint) ([]*entity.Production, error)
	FindById(db *gorm.DB, id uint) (*entity.Production, error)
	Create(db *gorm.DB, production *entity.Production) error
	Update(db *gorm.DB, production *entity.Production) error
	Delete(db *gorm.DB, id uint) error
}

type ProductionRepositoryImpl struct {
	Log *logrus.Logger
}

func NewProductionRepository(log *logrus.Logger) ProductionRepository {
	return &ProductionRepositoryImpl{
		Log: log,
	}
}

func (r *ProductionRepositoryImpl) FindAll(db *gorm.DB) ([]*entity.Production, error) {
	var productions []*entity.Production
	err := db.Order("created_at DESC").Find(&productions).Error
	return productions, err
}

func (r *ProductionRepositoryImpl) FindByMonth(db *gorm.DB, year uint, month uint) ([]*entity.Production, error) {
	var productions []*entity.Production
	err := db.
		Where("EXTRACT(YEAR FROM created_at) = ?", year).
		Where("EXTRACT(MONTH FROM created_at) = ?", month).
		Order("created_at DESC").
		Find(&productions).Error
	return productions, err
}

func (r *ProductionRepositoryImpl) FindById(db *gorm.DB, id uint) (*entity.Production, error) {
	var production entity.Production
	if err := db.First(&production, id).Error; err != nil {
		return nil, err
	}

	return &production, nil
}

func (r *ProductionRepositoryImpl) Create(db *gorm.DB, production *entity.Production) error {
	return db.Create(&production).Error
}

func (r *ProductionRepositoryImpl) Update(db *gorm.DB, production *entity.Production) error {
	return db.Save(&production).Error
}

func (r *ProductionRepositoryImpl) Delete(db *gorm.DB, id uint) error {
	var production entity.Production
	return db.Where("id = ?", id).Delete(&production).Error
}
