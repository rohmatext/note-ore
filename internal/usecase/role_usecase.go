package usecase

import (
	"context"
	"rohmatext/ore-note/internal/entity"
	"rohmatext/ore-note/internal/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RoleUseCase interface {
	GetAllRoles(ctx context.Context) ([]*entity.Role, error)
}

type RoleUseCaseImpl struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	RoleRepository repository.RoleRepository
}

func NewRoleUseCase(db *gorm.DB, log *logrus.Logger, roleRepo repository.RoleRepository) RoleUseCase {
	return &RoleUseCaseImpl{
		DB:             db,
		Log:            log,
		RoleRepository: roleRepo,
	}
}

func (uc *RoleUseCaseImpl) GetAllRoles(ctx context.Context) ([]*entity.Role, error) {
	db := uc.DB.WithContext(ctx)
	users, err := uc.RoleRepository.FindAll(db)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return users, nil
}
