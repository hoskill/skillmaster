package models

import "time"

type Material struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Title           string    `gorm:"not null" json:"title"`
	Description     string    `json:"description"`
	ContentURL      string    `gorm:"not null" json:"content_url"`
	DifficultyLevel string    `gorm:"check:difficulty_level IN ('beginner', 'intermediate', 'advanced')" json:"difficulty_level"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
