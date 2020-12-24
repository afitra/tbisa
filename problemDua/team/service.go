package team

type Service interface {
	RegisterTeam(input RegisterTeamInput) (Team, error)
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
