package usecase

import (
	"context"
	"rohmatext/ore-note/internal/entity"
	"rohmatext/ore-note/internal/model"
	"rohmatext/ore-note/internal/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductionUseCase interface {
	GetAllProductions(ctx context.Context) ([]*entity.Production, error)
	GetProductionById(ctx context.Context, id uint) (*entity.Production, error)
	CreateProduction(ctx context.Context, userId uint, request *model.CreateProductionRequest) (*entity.Production, error)
	UpdateProduction(ctx context.Context, userId uint, request *model.UpdateProductionRequest, id uint) (*entity.Production, error)
	DeleteProduction(ctx context.Context, id uint) error
}

type ProductionUseCaseImpl struct {
	DB                   *gorm.DB
	Log                  *logrus.Logger
	ProductionRepository repository.ProductionRepository
	UserRepository       repository.UserRepository
	OreRepository        repository.OreRepository
	SourceRepository     repository.SourceRepository
}

func NewProductionUseCase(db *gorm.DB, log *logrus.Logger, productionRepo repository.ProductionRepository, userRepo repository.UserRepository, oreRepo repository.OreRepository, sourceRepo repository.SourceRepository) ProductionUseCase {
	return &ProductionUseCaseImpl{
		DB:                   db,
		Log:                  log,
		ProductionRepository: productionRepo,
		UserRepository:       userRepo,
		OreRepository:        oreRepo,
		SourceRepository:     sourceRepo,
	}
}

func (uc *ProductionUseCaseImpl) GetAllProductions(ctx context.Context) ([]*entity.Production, error) {
	db := uc.DB.WithContext(ctx).Preload("User").Preload("Ore").Preload("Source")
	return uc.ProductionRepository.FindAll(db)
}

func (uc *ProductionUseCaseImpl) GetProductionById(ctx context.Context, id uint) (*entity.Production, error) {
	db := uc.DB.WithContext(ctx).Preload("User").Preload("Ore").Preload("Source")
	return uc.ProductionRepository.FindById(db, id)
}

func (uc *ProductionUseCaseImpl) CreateProduction(ctx context.Context, userId uint, request *model.CreateProductionRequest) (*entity.Production, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user, err := uc.UserRepository.FindById(tx, userId)
	if err != nil {
		uc.Log.Warnf("Failed to read user: %+v", err)
		return nil, err
	}

	ore, err := uc.OreRepository.FindById(tx, request.OreID)
	if err != nil {
		uc.Log.Warnf("Failed to read ore: %+v", err)
		return nil, err
	}

	source, err := uc.SourceRepository.FindById(tx, request.SourceID)
	if err != nil {
		uc.Log.Warnf("Failed to read source: %+v", err)
		return nil, err
	}

	production := entity.Production{
		UserID:   user.ID,
		OreID:    ore.ID,
		SourceID: source.ID,
		Weight:   request.Weight,
		Notes:    request.Notes,
	}

	if err := uc.ProductionRepository.Create(tx, &production); err != nil {
		uc.Log.Warnf("Failed to create production: %+v", err)
		return nil, err
	}

	preloadTx := tx.Preload("User").Preload("Ore").Preload("Source")
	result, err := uc.ProductionRepository.FindById(preloadTx, production.ID)
	if err != nil {
		uc.Log.Warnf("Failed to read production: %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.Warnf("Failed commit transaction: %+v", err)
		return nil, err
	}

	return result, nil
}

func (uc *ProductionUseCaseImpl) UpdateProduction(ctx context.Context, userId uint, request *model.UpdateProductionRequest, id uint) (*entity.Production, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	production, err := uc.ProductionRepository.FindById(tx, id)
	if err != nil {
		uc.Log.Warnf("Failed to read production data: %+v", err)
		return nil, err
	}

	user, err := uc.UserRepository.FindById(tx, userId)
	if err != nil {
		uc.Log.Warnf("Failed to read user: %+v", err)
		return nil, err
	}

	ore, err := uc.OreRepository.FindById(tx, request.OreID)
	if err != nil {
		uc.Log.Warnf("Failed to read ore: %+v", err)
		return nil, err
	}

	source, err := uc.SourceRepository.FindById(tx, request.SourceID)
	if err != nil {
		uc.Log.Warnf("Failed to read source: %+v", err)
		return nil, err
	}

	production.UserID = user.ID
	production.OreID = ore.ID
	production.SourceID = source.ID

	if err := uc.ProductionRepository.Update(tx, production); err != nil {
		uc.Log.Warnf("Failed to update production: %+v", err)
		return nil, err
	}

	preloadTx := tx.Preload("User").Preload("Ore").Preload("Source")
	result, err := uc.ProductionRepository.FindById(preloadTx, production.ID)
	if err != nil {
		uc.Log.Warnf("Failed to read production: %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.Warnf("Failed commit transaction: %+v", err)
		return nil, err
	}

	return result, nil
}

func (uc *ProductionUseCaseImpl) DeleteProduction(ctx context.Context, id uint) error {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := uc.ProductionRepository.Delete(tx, id); err != nil {
		uc.Log.Warnf("Failed to delete production: %+v", err)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.Warnf("Failed commit transaction: %+v", err)
		return err
	}

	return nil
}
