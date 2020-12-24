package main

import (
	"fmt"
	"log"

	"problemDua/handler"
	"problemDua/helper"
	"problemDua/player"
	"problemDua/team"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	databaseName := helper.GoDotEnvVariable("DATABASE_NAME")
	dsn := fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	db.Migrator().CreateTable(team.Team{})
	db.Migrator().CreateTable(player.Player{})

	fmt.Println("koneksi DB berhasil *******")

	playerRepository := player.NewRepository(db)
	playerService := player.NewService(playerRepository)

	teamRepository := team.NewRepository(db)
	teamService := team.NewService(teamRepository)

	playerHandler := handler.NewPlayerHandler(playerService)
	teamHandler := handler.NewTeamHandler(teamService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/player", playerHandler.RegisterPlayer)
	api.POST("/team", teamHandler.RegisterTeam)

	PORT := helper.GoDotEnvVariable("PORT")
	router.Run(fmt.Sprintf(":%s", PORT))

}
