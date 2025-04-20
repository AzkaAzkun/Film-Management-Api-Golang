package main

import (
	"film-management-api-golang/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to loading env file")
	}

	RestApi := config.NewRest()
	RestApi.Start()
}
