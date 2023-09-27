package main

import (
	"ki_assignment-1/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	router := gin.Default()

	config.SetupDatabaseConnection()

	router.Run()
}