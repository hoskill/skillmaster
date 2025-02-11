package repository

import (
	"context"
	"gorm.io/gorm"
	"skillmaster/pkg/entities/models"
)

type MaterialRepository interface {
	GetAll(ctx context.Context) ([]models.Material, error)
	GetByID(ctx context.Context, id uint) (*models.Material, error)
}

type materialRepository struct {
	db *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) MaterialRepository {
	return &materialRepository{db: db}
}

func (r *materialRepository) GetAll(ctx context.Context) ([]models.Material, error) {
	var materials []models.Material
	if err := r.db.WithContext(ctx).Find(&materials).Error; err != nil {
		return nil, err
	}
	return materials, nil
}

func (r *materialRepository) GetByID(ctx context.Context, id uint) (*models.Material, error) {
	var material models.Material
	if err := r.db.WithContext(ctx).First(&material, id).Error; err != nil {
		return nil, err
	}
	return &material, nil
}
