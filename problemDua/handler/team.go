package handler

import (
	"net/http"
	"problemDua/helper"
	"problemDua/team"

	"github.com/gin-gonic/gin"
)

type teamHandler struct {
	teamService team.Service
}

func NewTeamHandler(teamService team.Service) *teamHandler {
	return &teamHandler{teamService}
}

func (h *teamHandler) RegisterTeam(c *gin.Context) {

	var input team.RegisterTeamInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Register Team failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newTeam, err := h.teamService.RegisterTeam(input)
	if err != nil {
		response := helper.ApiResponse("Register Team failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := team.FormatTeam(newTeam)

	response := helper.ApiResponse("Team has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
