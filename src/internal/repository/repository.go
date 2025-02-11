package repository

import (
	"gorm.io/gorm"
)

type Repository interface {
	MaterialRepository
	AuthRepository
}

type repository struct {
	MaterialRepository
	AuthRepository
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{NewMaterialRepository(db), NewAuthRepository(db)}
}
