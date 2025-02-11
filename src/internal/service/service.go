package service

import (
	"skillmaster/internal/repository"
)

type Service interface {
	MaterialService
	AuthService
}

type service struct {
	AuthService
	MaterialService
}

func NewService(repo repository.Repository) Service {
	return &service{NewAuthService(repo), NewMaterialService(repo)}
}
