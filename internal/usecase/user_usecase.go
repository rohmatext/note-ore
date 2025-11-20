package usecase

import (
	"context"
	"rohmatext/ore-note/internal/entity"
	"rohmatext/ore-note/internal/model"
	"rohmatext/ore-note/internal/repository"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase interface {
	Login(ctx context.Context, request *model.LoginUserRequest) (*entity.User, error)
	GetUser(ctx context.Context, id uint) (*entity.User, error)
	GetUsers(ctx context.Context) ([]*entity.User, error)
}

type UserUseCaseImpl struct {
	Repository repository.UserRepository
	DB         *gorm.DB
	Log        *logrus.Logger
}

func NewUserUseCase(db *gorm.DB, log *logrus.Logger, userRepo repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{
		Repository: userRepo,
		DB:         db,
		Log:        log,
	}
}

func (uc *UserUseCaseImpl) Login(ctx context.Context, request *model.LoginUserRequest) (*entity.User, error) {
	user, err := uc.Repository.FindByUsername(uc.DB, request.Username)
	if err != nil {
		uc.Log.Warnf("Invalid find by username: %+v", err)
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		uc.Log.Warnf("Password comparison failed: %+v", err)
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

func (uc *UserUseCaseImpl) GetUser(ctx context.Context, id uint) (*entity.User, error) {
	db := uc.DB.WithContext(ctx)
	user, err := uc.Repository.FindById(db, id)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

func (uc *UserUseCaseImpl) GetUsers(ctx context.Context) ([]*entity.User, error) {
	db := uc.DB.WithContext(ctx)
	users, err := uc.Repository.FindAll(db)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return users, nil
}
