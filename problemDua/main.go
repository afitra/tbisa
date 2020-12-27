package main

import (
	"fmt"
	"log"

	"problemDua/handler"
	"problemDua/helper"
	"problemDua/team"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	databaseName := helper.GoDotEnvVariable("DATABASE_NAME")
	dsn := fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		log.Fatal(err.Error())

	}
	db.Migrator().CreateTable(team.Player{})
	db.Migrator().CreateTable(team.Team{})

	fmt.Println("koneksi DB berhasil *******")

	teamRepository := team.NewRepository(db)
	teamService := team.NewService(teamRepository)

	teamHandler := handler.NewTeamHandler(teamService)

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/player", teamHandler.CreatePlayer)
	api.POST("/player/add", teamHandler.AddPlayer)
	api.POST("/team", teamHandler.RegisterTeam)
	api.GET("/team/:id", teamHandler.GetTeam)

	PORT := helper.GoDotEnvVariable("PORT")
	router.Run(fmt.Sprintf(":%s", PORT))

}
