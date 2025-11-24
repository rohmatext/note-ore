package usecase

import (
	"context"
	"rohmatext/ore-note/internal/entity"
	"rohmatext/ore-note/internal/model"
	"rohmatext/ore-note/internal/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SourceUseCase interface {
	GetAllSources(ctx context.Context) ([]*entity.Source, error)
	GetSourceById(ctx context.Context, id uint) (*entity.Source, error)
	CreateSource(ctx context.Context, request *model.CreateSourceRequest) (*entity.Source, error)
	UpdateSource(ctx context.Context, request *model.UpdateSourceRequest, id uint) (*entity.Source, error)
	DeleteSource(ctx context.Context, id uint) error
}

type SourceUseCaseImpl struct {
	DB               *gorm.DB
	Log              *logrus.Logger
	SourceRepository repository.SourceRepository
}

func NewSourceUseCase(db *gorm.DB, log *logrus.Logger, sourceRepo repository.SourceRepository) SourceUseCase {
	return &SourceUseCaseImpl{
		DB:               db,
		Log:              log,
		SourceRepository: sourceRepo,
	}
}

func (uc *SourceUseCaseImpl) GetAllSources(ctx context.Context) ([]*entity.Source, error) {
	db := uc.DB.WithContext(ctx)
	return uc.SourceRepository.FindAll(db)
}

func (uc *SourceUseCaseImpl) GetSourceById(ctx context.Context, id uint) (*entity.Source, error) {
	db := uc.DB.WithContext(ctx)
	return uc.SourceRepository.FindById(db, id)
}

func (uc *SourceUseCaseImpl) CreateSource(ctx context.Context, request *model.CreateSourceRequest) (*entity.Source, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	source := entity.Source{
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
	}

	if err := uc.SourceRepository.Create(tx, &source); err != nil {
		uc.Log.Warnf("Failed to create source: %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.Warnf("Failed commit transaction: %+v", err)
		return nil, err
	}

	return &source, nil
}

func (uc *SourceUseCaseImpl) UpdateSource(ctx context.Context, request *model.UpdateSourceRequest, id uint) (*entity.Source, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	source, err := uc.SourceRepository.FindById(tx, id)
	if err != nil {
		uc.Log.Warnf("Failed to read source data: %+v", err)
		return nil, err
	}

	source.Name = request.Name
	source.PhoneNumber = request.PhoneNumber
	if err := uc.SourceRepository.Update(tx, source); err != nil {
		uc.Log.Warnf("Failed to update source: %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.Warnf("Failed commit transaction: %+v", err)
		return nil, err
	}

	return source, nil
}

func (uc *SourceUseCaseImpl) DeleteSource(ctx context.Context, id uint) error {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := uc.SourceRepository.Delete(tx, id); err != nil {
		uc.Log.Warnf("Failed to delete source: %+v", err)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.Warnf("Failed commit transaction: %+v", err)
		return err
	}

	return nil
}
