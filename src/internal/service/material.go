package service

import (
	"context"
	"skillmaster/internal/repository"
	"skillmaster/pkg/entities/models"
)

type MaterialService interface {
	GetMaterials(ctx context.Context) ([]models.Material, error)
	GetMaterial(ctx context.Context, id uint) (*models.Material, error)
}

type materialService struct {
	repo repository.MaterialRepository
}

func NewMaterialService(repo repository.MaterialRepository) MaterialService {
	return &materialService{repo: repo}
}

func (s *materialService) GetMaterials(ctx context.Context) ([]models.Material, error) {
	return s.repo.GetAll(ctx)
}

func (s *materialService) GetMaterial(ctx context.Context, id uint) (*models.Material, error) {
	return s.repo.GetByID(ctx, id)
}
