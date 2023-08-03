package initializer

import (
	"fmt"

	env "github.com/joho/godotenv"
)

func LoadEnvVar() {
	err := env.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
