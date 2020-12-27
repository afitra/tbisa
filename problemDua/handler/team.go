package handler

import (
	"fmt"
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

		response := helper.ApiResponse("Register Team failed, error input", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
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

func (h *teamHandler) GetTeam(c *gin.Context) {

	var input team.GetTeamInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.ApiResponse("Failed to get  Team, id Not FOund", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	detailTeam, err := h.teamService.GetTeamByID(input.ID)
	fmt.Println(">>>>", detailTeam)
	if err != nil {
		response := helper.ApiResponse("Failed to get Team", http.StatusBadRequest, "error", err)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Team Detail", http.StatusOK, "success", team.FormatTeamDetail(detailTeam))

	c.JSON(http.StatusOK, response)
}

func (h *teamHandler) CreatePlayer(c *gin.Context) {

	var input team.GetCreatePlayerInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.ApiResponse("Failed to Register  Player, input ", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	newPlayer, err := h.teamService.CreatePlayer(input)

	if err != nil {
		response := helper.ApiResponse("Failed to Register Player", http.StatusBadRequest, "error", err)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Player Register", http.StatusOK, "success", team.FormatPlayer(newPlayer))

	c.JSON(http.StatusOK, response)

}

func (h *teamHandler) AddPlayer(c *gin.Context) {

	var input team.AddPlayerInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.ApiResponse("Failed to add  Player to Teammm ", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	data, err := h.teamService.AddPlayer(input)

	if err != nil {
		response := helper.ApiResponse("Failed to Register Player", http.StatusBadRequest, "error", err)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Player Register", http.StatusOK, "success", team.FormatPlayer(data))

	c.JSON(http.StatusOK, response)

}
