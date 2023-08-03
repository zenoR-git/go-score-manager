package main

import (
	initializer "score-manager/initializers"
	model "score-manager/models"
)

func init() {
	initializer.LoadEnvVar()
	initializer.ConnectDB()
}

func main() {
	initializer.DB.AutoMigrate(&model.UserScore{})
}
