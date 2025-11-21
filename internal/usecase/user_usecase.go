package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
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
	Login(ctx context.Context, request *model.LoginUserRequest) (*model.TokenPair, error)
	GetUserById(ctx context.Context, id uint) (*entity.User, error)
	GetAllUsers(ctx context.Context) ([]*entity.User, error)
	RefreshAccessToken(ctx context.Context, refreshToken string) (*model.TokenPair, error)
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

func (uc *UserUseCaseImpl) Login(ctx context.Context, request *model.LoginUserRequest) (*model.TokenPair, error) {
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

	accessToken, plainToken, refreshToken, err := uc.generateTokens(user.ID)
	if err != nil {
		uc.Log.Warnf("Failed to generate tokens: %+v", err)
		return nil, ErrInvalidToken
	}

	if err := uc.RefreshTokenRepository.Create(tx, refreshToken); err != nil {
		uc.Log.Warnf("Failed save refresh token: %+v", err)
		return nil, ErrInvalidCredentials
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.Warnf("Failed commit transaction: %+v", err)
		return nil, ErrInvalidCredentials
	}

	return &model.TokenPair{
		AccessToken: *accessToken,
		RefreshToken: model.RefreshTokenResponse{
			Token:     *plainToken,
			ExpiresAt: refreshToken.ExpiredAt,
		},
	}, nil
}

func (uc *UserUseCaseImpl) GetUserById(ctx context.Context, id uint) (*entity.User, error) {
	db := uc.DB.WithContext(ctx)
	user, err := uc.UserRepository.FindById(db, id)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

func (uc *UserUseCaseImpl) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	db := uc.DB.WithContext(ctx)
	users, err := uc.UserRepository.FindAll(db)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return users, nil
}

func (uc *UserUseCaseImpl) RefreshAccessToken(ctx context.Context, refreshToken string) (*model.TokenPair, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	hash := sha256.Sum256([]byte(refreshToken))
	tokenStr := hex.EncodeToString(hash[:])

	token, err := uc.RefreshTokenRepository.FindByToken(tx, tokenStr)
	if err != nil {
		return nil, ErrInvalidToken
	}

	accessToken, plainToken, newRefreshToken, err := uc.generateTokens(token.UserID)
	if err != nil {
		return nil, ErrInvalidToken
	}

	if err := uc.RefreshTokenRepository.Create(tx, newRefreshToken); err != nil {
		uc.Log.Warnf("Failed create refresh token: %+v", err)
		return nil, ErrInvalidToken
	}

	if err := uc.RefreshTokenRepository.Delete(tx, token); err != nil {
		return nil, ErrInvalidToken
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.Warnf("Failed coomit transaction: %+v", err)
		return nil, ErrInvalidToken
	}

	return &model.TokenPair{
		AccessToken: *accessToken,
		RefreshToken: model.RefreshTokenResponse{
			Token:     *plainToken,
			ExpiresAt: newRefreshToken.ExpiredAt,
		},
	}, nil
}

func (uc *UserUseCaseImpl) generateTokens(userId uint) (*string, *string, *entity.RefreshToken, error) {
	// generate access token
	accessToken, err := uc.TokenService.GenerateToken(uint(userId))
	if err != nil {
		uc.Log.Warnf("Failed generate access token: %+v", err)
		return nil, nil, nil, ErrInvalidToken
	}

	// generate refresh token
	plainRefreshToken, err := stringx.Random(20)
	hash := sha256.Sum256([]byte(plainRefreshToken))
	token := hex.EncodeToString(hash[:])

	newRefreshToken := *new(entity.RefreshToken)
	newRefreshToken.Token = token
	newRefreshToken.UserID = userId
	newRefreshToken.ExpiredAt = time.Now().Add(7 * 24 * time.Hour)

	if err != nil {
		uc.Log.Warnf("Failed create random string: %+v", err)
		return nil, nil, nil, ErrInvalidToken
	}

	return &accessToken, &plainRefreshToken, &newRefreshToken, nil
}
