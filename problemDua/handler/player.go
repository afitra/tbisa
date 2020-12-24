package handler

import (
	"net/http"
	"problemDua/helper"
	"problemDua/player"

	"github.com/gin-gonic/gin"
)

type playerHandler struct {
	playerService player.Service
}

func NewPlayerHandler(playerService player.Service) *playerHandler {
	return &playerHandler{playerService}
}

func (h *playerHandler) RegisterPlayer(c *gin.Context) {

	var input player.RegisterPlayerInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Register Player failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPlayer, err := h.playerService.RegisterPlayer(input)
	if err != nil {
		response := helper.ApiResponse("Register Team failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := player.FormatPlayer(newPlayer)

	response := helper.ApiResponse("Player has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
