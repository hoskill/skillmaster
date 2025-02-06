package service

import (
	"skillmaster/internal/repository"
)

type Service interface {
	MaterialService
	AuthService
}

type service struct {
	repo repository.Repository
	MaterialService
	AuthService
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}
