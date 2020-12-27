package team

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(team Team) (Team, error)
	FindTeamById(id int) (Team, error)
	FindPlayerById(id int) (Player, error)
	CreatePlayer(input Player) (Player, error)
	UpdatePlayer(input Player) (Player, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}

}

func (r *repository) Save(team Team) (Team, error) {
	err := r.db.Create(&team).Error

	if err != nil {
		return team, err
	}
	return team, nil
}

func (r *repository) FindTeamById(id int) (Team, error) {
	var team Team

	err := r.db.Preload("Player").Where("id = ?", id).Find(&team).Error

	if err != nil {

		return team, err
	}

	return team, nil
}

func (r *repository) FindPlayerById(id int) (Player, error) {
	var player Player

	err := r.db.Where("id = ?", id).Find(&player).Error

	if err != nil {

		return player, err
	}

	return player, nil
}

func (r *repository) CreatePlayer(input Player) (Player, error) {

	err := r.db.Create(&input).Error

	if err != nil {
		return input, err
	}
	return input, nil

}

func (r *repository) UpdatePlayer(input Player) (Player, error) {

	err := r.db.Model(&Player{}).Where("id = ?", input.ID).Update("team_id", input.TeamID).Error

	if err != nil {
		return input, err
	}

	return input, nil

}
