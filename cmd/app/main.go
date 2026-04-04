package main

import (
	"film-management-api-golang/internal/config"
	"time"
	_ "time/tzdata"

	"github.com/joho/godotenv"
)

func main() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic("Failed to load timezone: " + err.Error())
	}
	time.Local = loc

	if err := godotenv.Load(); err != nil {
		panic("Failed to loading env file")
	}

	RestApi := config.NewRest()
	RestApi.Start()
}
