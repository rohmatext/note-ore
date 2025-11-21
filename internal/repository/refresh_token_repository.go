package repository

import (
	"rohmatext/ore-note/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	Create(db *gorm.DB, token *entity.RefreshToken) error
}

type RefreshTokenRepositoryImpl struct {
	Log *logrus.Logger
}

func NewRefreshTokenRepository(log *logrus.Logger) RefreshTokenRepository {
	return &RefreshTokenRepositoryImpl{
		Log: log,
	}
}

func (r *RefreshTokenRepositoryImpl) Create(db *gorm.DB, token *entity.RefreshToken) error {
	if err := db.Create(token).Error; err != nil {
		return err
	}

	return nil
}
