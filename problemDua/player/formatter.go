package player

import "problemDua/team"

type PlayerFormatter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	TeamID int    `json:"team_id"`
}

type PlayerDetailFormatter struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Age    int       `json:"age"`
	TeamID int       `json:"team_id"`
	Team   team.Team `json:"team"`
}

func FormatPlayer(player Player) PlayerFormatter {

	playerFormatter := PlayerFormatter{}
	playerFormatter.ID = player.ID
	playerFormatter.Name = player.Name
	playerFormatter.Age = player.Age
	playerFormatter.TeamID = player.TeamID

	return playerFormatter
}

func FormatDetailPlayer(player Player) PlayerDetailFormatter {

	playerFormatter := PlayerDetailFormatter{}
	playerFormatter.ID = player.ID
	playerFormatter.Name = player.Name
	playerFormatter.Age = player.Age
	playerFormatter.TeamID = player.TeamID

	return playerFormatter
}

func FormatPlayers(players []Player) []PlayerDetailFormatter {

	allFormatter := []PlayerDetailFormatter{}
	for _, player := range players {

		playerFormatter := FormatDetailPlayer(player)

		allFormatter = append(allFormatter, playerFormatter)
	}

	return allFormatter
}
