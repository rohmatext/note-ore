package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"rohmatext/ore-note/internal/entity"
	"rohmatext/ore-note/internal/model"
	"rohmatext/ore-note/internal/repository"
	"rohmatext/ore-note/internal/utils/stringx"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase interface {
	Login(ctx context.Context, request *model.LoginUserRequest) (*model.LoginOutput, error)
	GetUser(ctx context.Context, id uint) (*entity.User, error)
	GetUsers(ctx context.Context) ([]*entity.User, error)
}

type UserUseCaseImpl struct {
	DB                     *gorm.DB
	Log                    *logrus.Logger
	UserRepository         repository.UserRepository
	RefreshTokenRepository repository.RefreshTokenRepository
	TokenService           TokenService
}

func NewUserUseCase(db *gorm.DB, log *logrus.Logger, refreshTokenRepo repository.RefreshTokenRepository, userRepo repository.UserRepository, token TokenService) UserUseCase {
	return &UserUseCaseImpl{
		DB:                     db,
		Log:                    log,
		UserRepository:         userRepo,
		RefreshTokenRepository: refreshTokenRepo,
		TokenService:           token,
	}
}

func (uc *UserUseCaseImpl) Login(ctx context.Context, request *model.LoginUserRequest) (*model.LoginOutput, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user, err := uc.UserRepository.FindByUsername(tx, request.Username)
	if err != nil {
		uc.Log.Warnf("Invalid find by username: %+v", err)
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		uc.Log.Warnf("Password comparison failed: %+v", err)
		return nil, ErrInvalidCredentials
	}

	accessToken, err := uc.TokenService.GenerateToken(uint(user.ID))
	if err != nil {
		uc.Log.Warnf("Failed generate access token: %+v", err)
		return nil, ErrInvalidCredentials
	}

	plainToken, err := stringx.Random(20)
	if err != nil {
		uc.Log.Warnf("Failed create random string: %+v", err)
		return nil, ErrInvalidCredentials
	}

	hash := sha256.Sum256([]byte(plainToken))
	token := hex.EncodeToString(hash[:])
	refreshToken := entity.RefreshToken{Token: token, UserID: user.ID, ExpiredAt: time.Now().Add(7 * 24 * time.Hour)}
	if err := uc.RefreshTokenRepository.Create(tx, &refreshToken); err != nil {
		uc.Log.Warnf("Failed create refresh token: %+v", err)
		return nil, ErrInvalidCredentials
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.Warnf("Failed coomit transaction: %+v", err)
		return nil, ErrInvalidCredentials
	}

	return &model.LoginOutput{
		AccessToken: accessToken,
		RefreshToken: model.RefreshTokenResponse{
			Token:     fmt.Sprintf("%d|%s", refreshToken.ID, plainToken),
			ExpiresAt: refreshToken.ExpiredAt,
		},
	}, nil
}

func (uc *UserUseCaseImpl) GetUser(ctx context.Context, id uint) (*entity.User, error) {
	db := uc.DB.WithContext(ctx)
	user, err := uc.UserRepository.FindById(db, id)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

func (uc *UserUseCaseImpl) GetUsers(ctx context.Context) ([]*entity.User, error) {
	db := uc.DB.WithContext(ctx)
	users, err := uc.UserRepository.FindAll(db)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return users, nil
}
