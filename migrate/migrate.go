package main

import (
	"fmt"
	"log"

	"user-service/initializers"
	"user-service/models"
)

func init() {
	config, err := initializers.LoadFromEnv()
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
