package player

import (
	"problemDua/team"
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	Name      string
	Age       int
	TeamID    int
	Team      team.Team `gorm:"foreignKey:TeamID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
