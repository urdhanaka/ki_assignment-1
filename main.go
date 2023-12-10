package main

import (
	"ki_assignment-1/config"
	"ki_assignment-1/controllers"
	"ki_assignment-1/middleware"
	"ki_assignment-1/repository"
	"ki_assignment-1/routes"
	"ki_assignment-1/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	var (
		db = config.SetupDatabaseConnection()

		jwtService = service.NewJWTService()

		userRepository = repository.NewUserRepository(db)
		fileRepository = repository.NewFileRepository(db)

		userService = service.NewUserService(userRepository)
		fileService = service.NewFileService(fileRepository, userRepository)

		userController = controllers.NewUserController(userService, jwtService)
		fileController = controllers.NewFileController(fileService, jwtService, userService)
	)

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	routes.UserRoutes(router, userController, jwtService)
	routes.FileRoutes(router, fileController, jwtService)

	router.Run()
}
