package repository

import (
	"gorm.io/gorm"
)

type Repository interface {
	MaterialRepository
	AuthRepository
}

type repository struct {
	db *gorm.DB
	MaterialRepository
	AuthRepository
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
