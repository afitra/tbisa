package team

type TeamFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type TeamDetailFormatter struct {
	ID     int               `json:"id"`
	Name   string            `json:"name"`
	Player []PlayerFormatter `json:"players"`
}

type PlayerFormatter struct {
	ID       int    `json:"id"`
	TeamID   int    `json:"team_id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

type AddPlayerInput struct {
	TeamID   int `json:"team_id"`
	PlayerID int `json:"player_id"`
}

func FormatTeam(team Team) TeamFormatter {

	return TeamFormatter{
		ID:   team.ID,
		Name: team.Name,
	}

}

func FormatTeamDetail(team Team) TeamDetailFormatter {

	data := TeamDetailFormatter{
		ID:     team.ID,
		Name:   team.Name,
		Player: FormatMultiPlayer(team.Player),
	}

	return data
}

func FormatMultiPlayer(players []Player) []PlayerFormatter {
	var result []PlayerFormatter
	for _, item := range players {
		result = append(result, FormatPlayer(item))
	}

	return result
}
func FormatPlayer(player Player) PlayerFormatter {
	return PlayerFormatter{
		ID:       player.ID,
		TeamID:   player.TeamID,
		Name:     player.Name,
		Age:      player.Age,
		Position: player.Position,
	}
}
