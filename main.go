package main

import (
	"score-manager/controller"
	initializer "score-manager/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVar()
	initializer.ConnectDB()
}

func main() {
	r := gin.Default()

	r.POST("/players", controller.CreatePlayerController)
	r.PUT("/players/:id", controller.UpdatePlayerController)
	r.DELETE("/players/:id", controller.DeletePlayer)
	r.GET("/players", controller.GetAllPlayers)
	r.GET("/players/rank/:val", controller.GetRankedPlayer)
	r.GET("/players/random", controller.GetRandom)

	r.Run(":8080")

}
