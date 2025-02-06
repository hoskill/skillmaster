package entities

import (
	"time"
)

type UserMaterialProgress struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	UserID             uint      `gorm:"not null" json:"user_id"`
	MaterialID         uint      `gorm:"not null" json:"material_id"`
	IsCompleted        bool      `gorm:"default:false" json:"is_completed"`
	ProgressPercentage int       `gorm:"check:progress_percentage BETWEEN 0 AND 100" json:"progress_percentage"`
	LastAccessed       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"last_accessed"`
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
