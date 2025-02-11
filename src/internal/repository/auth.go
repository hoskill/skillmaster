package repository

import (
	"context"
	"gorm.io/gorm/clause"
	"skillmaster/pkg/entities/models"
	"time"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserByUserID(ctx context.Context, id string) (*models.User, error)

	SaveRefreshToken(ctx context.Context, userID, token string, expiresAt time.Time) error
	GetRefreshToken(ctx context.Context, token string) (*models.RefreshToken, error)
	DeleteRefreshToken(ctx context.Context, token string) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *authRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) GetUserByUserID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) SaveRefreshToken(ctx context.Context, userID, token string, expiresAt time.Time) error {
	refreshToken := models.RefreshToken{
		Token:     token,
		UserID:    userID,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
	}

	return r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"token", "expires_at"}),
	}).Create(&refreshToken).Error
}

func (r *authRepository) GetRefreshToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	if err := r.db.WithContext(ctx).Where("token = ?", token).First(&refreshToken).Error; err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

func (r *authRepository) DeleteRefreshToken(ctx context.Context, token string) error {
	return r.db.WithContext(ctx).Where("token = ?", token).Delete(&models.RefreshToken{}).Error
}
