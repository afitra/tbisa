package team

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	ID        int
	Name      string
	Player    []Player
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type Team
type Player struct {
	ID        int
	TeamID    int
	Name      string
	Age       int
	Position  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
