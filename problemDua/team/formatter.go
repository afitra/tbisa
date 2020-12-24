package team

type TeamFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatTeam(team Team) TeamFormatter {

	formatter := TeamFormatter{
		ID:   team.ID,
		Name: team.Name,
	}
	return formatter
}
