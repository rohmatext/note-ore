package usecase

import (
	"context"
	"rohmatext/ore-note/internal/entity"
	"rohmatext/ore-note/internal/model"
	"rohmatext/ore-note/internal/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OreUseCase interface {
	GetAllOres(ctx context.Context) ([]*entity.Ore, error)
	GetOreById(ctx context.Context, id uint16) (*entity.Ore, error)
	CreateOre(ctx context.Context, request *model.CreateOreRequest) (*entity.Ore, error)
	UpdateOre(ctx context.Context, request *model.UpdateOreRequest, id uint16) (*entity.Ore, error)
	DeleteOre(ctx context.Context, id uint16) error
}

type OreUseCaseImpl struct {
	DB            *gorm.DB
	Log           *logrus.Logger
	OreRepository repository.OreRepository
}

func NewOreUseCase(db *gorm.DB, log *logrus.Logger, oreRepo repository.OreRepository) OreUseCase {
	return &OreUseCaseImpl{
		DB:            db,
		Log:           log,
		OreRepository: oreRepo,
	}
}

func (uc *OreUseCaseImpl) GetAllOres(ctx context.Context) ([]*entity.Ore, error) {
	db := uc.DB.WithContext(ctx)
	return uc.OreRepository.FindAll(db)
}

func (uc *OreUseCaseImpl) GetOreById(ctx context.Context, id uint16) (*entity.Ore, error) {
	db := uc.DB.WithContext(ctx)
	return uc.OreRepository.FindById(db, id)
}

func (uc *OreUseCaseImpl) CreateOre(ctx context.Context, request *model.CreateOreRequest) (*entity.Ore, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	ore := entity.Ore{
		Name: request.Name,
	}

	if err := uc.OreRepository.Create(tx, &ore); err != nil {
		uc.Log.Warnf("Failed to create ore: %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.Warnf("Failed commit transaction: %+v", err)
		return nil, err
	}

	return &ore, nil
}

func (uc *OreUseCaseImpl) UpdateOre(ctx context.Context, request *model.UpdateOreRequest, id uint16) (*entity.Ore, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	ore, err := uc.OreRepository.FindById(tx, id)
	if err != nil {
		uc.Log.Warnf("Failed to read ore data: %+v", err)
		return nil, err
	}

	ore.Name = request.Name
	if err := uc.OreRepository.Update(tx, ore); err != nil {
		uc.Log.Warnf("Failed to update ore: %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.Warnf("Failed commit transaction: %+v", err)
		return nil, err
	}

	return ore, nil
}

func (uc *OreUseCaseImpl) DeleteOre(ctx context.Context, id uint16) error {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := uc.OreRepository.Delete(tx, id); err != nil {
		uc.Log.Warnf("Failed to delete ore: %+v", err)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.Warnf("Failed commit transaction: %+v", err)
		return err
	}

	return nil
}
