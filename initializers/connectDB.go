package initializer

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	//dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + databaseName + " port=" + port + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in connecting to database")
	}
}
