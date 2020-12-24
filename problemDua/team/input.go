package team

type RegisterTeamInput struct {
	Name string `json:"name" binding:"required"`
}
