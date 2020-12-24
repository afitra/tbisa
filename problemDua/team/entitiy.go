package team

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
