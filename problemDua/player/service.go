package player

type Service interface {
	RegisterPlayer(input RegisterPlayerInput) (Player, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterPlayer(input RegisterPlayerInput) (Player, error) {

	player := Player{}
	player.Name = input.Name
	player.Age = input.Age
	player.TeamID = input.TeamID

	newPlayer, err := s.repository.Save(player)
	if err != nil {
		return newPlayer, err
	}

	return newPlayer, nil

}
