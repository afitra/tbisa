package player

type GetPlayerDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type RegisterPlayerInput struct {
	Name   string `json:"name" binding:"required"`
	Age    int    `json:"age" binding:"required"`
	TeamID int    `json:"team_id" binding:"required"`
}
