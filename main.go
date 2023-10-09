package main

import (
	"ki_assignment-1/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := config.SetupDatabaseConnection()
	config.CloseDatabaseConnection(db)

	router := gin.Default()

	router.Run(":8080")
}