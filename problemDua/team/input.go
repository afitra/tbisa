package team

type RegisterTeamInput struct {
	Name string `json:"name" binding:"required"`
}

type GetTeamInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetCreatePlayerInput struct {
	Name     string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Position string `json:"position" binding:"required"`
}
