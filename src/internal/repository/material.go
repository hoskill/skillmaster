package repository

import (
	"context"
	"gorm.io/gorm"
	"skillmaster/pkg/entities"
)

type MaterialRepository interface {
	GetAll(ctx context.Context) ([]entities.Material, error)
	GetByID(ctx context.Context, id uint) (*entities.Material, error)
}

type materialRepository struct {
	db *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) MaterialRepository {
	return &materialRepository{db: db}
}

func (r *materialRepository) GetAll(ctx context.Context) ([]entities.Material, error) {
	var materials []entities.Material
	if err := r.db.WithContext(ctx).Find(&materials).Error; err != nil {
		return nil, err
	}
	return materials, nil
}

func (r *materialRepository) GetByID(ctx context.Context, id uint) (*entities.Material, error) {
	var material entities.Material
	if err := r.db.WithContext(ctx).First(&material, id).Error; err != nil {
		return nil, err
	}
	return &material, nil
}
