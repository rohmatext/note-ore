package repository

import (
	"rohmatext/ore-note/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(db *gorm.DB) ([]*entity.User, error)
	FindById(db *gorm.DB, id uint) (*entity.User, error)
	FindByUsername(db *gorm.DB, username string) (*entity.User, error)
	Create(db *gorm.DB, user *entity.User) error
	Update(db *gorm.DB, user *entity.User) error
}

type UserRepositoryImpl struct {
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) UserRepository {
	return &UserRepositoryImpl{
		Log: log,
	}
}

func (r *UserRepositoryImpl) FindAll(db *gorm.DB) ([]*entity.User, error) {
	var users []*entity.User
	err := db.Order("created_at DESC").Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) FindByUsername(db *gorm.DB, username string) (*entity.User, error) {
	var user entity.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindById(db *gorm.DB, id uint) (*entity.User, error) {
	var user entity.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Create(db *gorm.DB, user *entity.User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) Update(db *gorm.DB, user *entity.User) error {
	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
