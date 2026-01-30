package models

import "time"

type RefreshToken struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	TokenHash string
	ExpiresAt time.Time
}
