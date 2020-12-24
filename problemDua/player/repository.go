package player

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(player Player) (Player, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}

}

func (r *repository) Save(player Player) (Player, error) {
	err := r.db.Create(&player).Error

	if err != nil {
		return player, err
	}

	return player, nil
}
