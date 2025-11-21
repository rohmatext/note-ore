package usecase

import (
	"context"
	"rohmatext/ore-note/internal/entity"
	"rohmatext/ore-note/internal/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RefreshTokenUseCase interface {
	CreateToken(ctx context.Context, userId uint) (*entity.RefreshToken, error)
}

type RefreshTokenUseCaseImpl struct {
	Repository repository.RefreshTokenRepository
	DB         *gorm.DB
	Log        *logrus.Logger
}

func NewRefreshTokenUseCase(db *gorm.DB, log *logrus.Logger, refreshTokenRepository repository.RefreshTokenRepository) RefreshTokenUseCase {
	return &RefreshTokenUseCaseImpl{
		Repository: refreshTokenRepository,
		DB:         db,
		Log:        log,
	}
}

func (uc *RefreshTokenUseCaseImpl) CreateToken(ctx context.Context, userID uint) (*entity.RefreshToken, error) {
	return new(entity.RefreshToken), nil
}
