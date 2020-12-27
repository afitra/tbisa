package team

import (
	"errors"
)

type Service interface {
	RegisterTeam(input RegisterTeamInput) (Team, error)
	GetTeamByID(id int) (Team, error)
	CreatePlayer(input GetCreatePlayerInput) (Player, error)
	AddPlayer(input AddPlayerInput) (Player, error)
}

type service struct {
	repository Repository // repo yg hutruf kecil itu variable punya service sendiri bukan punya Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterTeam(input RegisterTeamInput) (Team, error) {

	team := Team{}
	team.Name = input.Name

	newTeam, err := s.repository.Save(team)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil
}

func (s *service) GetTeamByID(id int) (Team, error) {

	detailTeam, err := s.repository.FindTeamById(id)

	if err != nil {

		return detailTeam, err
	}
	return detailTeam, nil
}

func (s *service) CreatePlayer(input GetCreatePlayerInput) (Player, error) {

	data := Player{
		Name:     input.Name,
		Age:      input.Age,
		Position: input.Position,
	}

	newPlayer, err := s.repository.CreatePlayer(data)

	if err != nil {

		return newPlayer, err
	}
	return newPlayer, nil

}

func (s *service) AddPlayer(input AddPlayerInput) (Player, error) {

	player, err := s.repository.FindPlayerById(input.PlayerID)

	if err != nil {

		return player, err

	}

	if player.ID == 0 {
		return player, errors.New("Player ID not FOund")

	}

	team, err := s.repository.FindTeamById(input.TeamID)

	if err != nil {

		return player, err

	}

	if team.ID == 0 {
		return player, errors.New("Team ID not FOund")

	}

	player.TeamID = team.ID

	_, err = s.repository.UpdatePlayer(player)

	if err != nil {
		return player, err
	}

	newData, err := s.repository.FindPlayerById(input.PlayerID)

	if err != nil {
		return newData, err
	}

	return newData, nil

}
