package repository

import (
	"rohmatext/ore-note/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAll(db *gorm.DB) ([]*entity.Role, error)
	FindByName(db *gorm.DB, name string) (*entity.Role, error)
}

type RoleRepositoryImpl struct {
	Log *logrus.Logger
}

func NewRoleRepository(log *logrus.Logger) RoleRepository {
	return &RoleRepositoryImpl{
		Log: log,
	}
}

func (r *RoleRepositoryImpl) FindAll(db *gorm.DB) ([]*entity.Role, error) {
	var roles []*entity.Role
	err := db.Find(&roles).Error
	return roles, err
}

func (r *RoleRepositoryImpl) FindByName(db *gorm.DB, name string) (*entity.Role, error) {
	var role *entity.Role
	if err := db.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}

	return role, nil
}
