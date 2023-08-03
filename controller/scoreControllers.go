package controller

import (
	"fmt"
	initializer "score-manager/initializers"
	model "score-manager/models"
	"strconv"

	"math/rand"

	"github.com/gin-gonic/gin"
)

func CreatePlayerController(c *gin.Context) {
	var CreatePlayerReq struct {
		Name    string `json:"name"`
		Country string `json:"country"`
		Score   int    `json:"score"`
	}

	c.Bind(&CreatePlayerReq)

	if len(CreatePlayerReq.Name) > 15 {
		c.JSON(400, gin.H{
			"message": "name should atmost be 15 characters",
		})
		return
	}

	post := model.UserScore{Name: CreatePlayerReq.Name, Country: CreatePlayerReq.Country, Score: CreatePlayerReq.Score}
	result := initializer.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePlayerController(c *gin.Context) {
	var UpdatePlayer struct {
		Name  string `json:"name"`
		Score int    `json:"score"`
	}

	c.Bind(&UpdatePlayer)

	playerId := c.Param("id")

	var player model.UserScore

	//check if the player with that id exists
	result := initializer.DB.First(&player, playerId)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Id",
			"error":   result.Error,
		})
	}

	//update the player data
	player.Name = UpdatePlayer.Name
	player.Score = UpdatePlayer.Score
	initializer.DB.Save(&player)
	c.JSON(200, gin.H{
		"message": "data updated successfully",
	})

}

func DeletePlayer(c *gin.Context) {

	playerId := c.Param("id")

	var player model.UserScore

	//check if the player with that id exists
	result := initializer.DB.First(&player, playerId)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Id",
			"error":   result.Error,
		})
	}

	//delete the user
	initializer.DB.Delete(&model.UserScore{}, playerId)

	c.JSON(200, gin.H{
		"message": "data deleted successfully",
	})
}

func GetAllPlayers(c *gin.Context) {

	var players []model.UserScore

	result := initializer.DB.Order("score desc").Find(&players)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "error in retrieving data",
			"error":   result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"data": players,
	})

}

func GetRankedPlayer(c *gin.Context) {

	playerRank, err := strconv.Atoi(c.Param("val"))
	if err != nil {
		c.Status(400)
		fmt.Println("error in converting string to number")
		return
	}

	//if rank value is 0 or less
	if playerRank <= 0 {
		c.JSON(400, gin.H{
			"message": "rank should be greater than 0",
		})
		return
	}

	var players []model.UserScore

	result := initializer.DB.Order("score desc").Limit(1).Offset(playerRank - 1).Find(&players)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "error in retrieving data",
			"error":   result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"data": players,
	})
}

func GetRandom(c *gin.Context) {

	var count int64
	var players []model.UserScore

	result := initializer.DB.Find(&players).Count(&count)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "error in retrieving data",
			"error":   result.Error,
		})
		return
	}

	index := rand.Intn(int(count - 0))
	fmt.Println(index)
	c.JSON(200, gin.H{
		"data": players[index],
	})

}
