package repository

import (
	"context"
	"skillmaster/pkg/entities"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUserByUsername(ctx context.Context, username string) (*entities.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateUser(ctx context.Context, user *entities.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *authRepository) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
