package service

import (
	"context"
	"skillmaster/internal/repository"
	"skillmaster/pkg/entities"
)

type MaterialService interface {
	GetMaterials(ctx context.Context) ([]entities.Material, error)
	GetMaterial(ctx context.Context, id uint) (*entities.Material, error)
}

type materialService struct {
	repo repository.MaterialRepository
}

func NewMaterialService(repo repository.MaterialRepository) MaterialService {
	return &materialService{repo: repo}
}

func (s *materialService) GetMaterials(ctx context.Context) ([]entities.Material, error) {
	return s.repo.GetAll(ctx)
}

func (s *materialService) GetMaterial(ctx context.Context, id uint) (*entities.Material, error) {
	return s.repo.GetByID(ctx, id)
}
