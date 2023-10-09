package main

import (
	"ki_assignment-1/config"
	"ki_assignment-1/controllers"
	"ki_assignment-1/repository"
	"ki_assignment-1/routes"
	"ki_assignment-1/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}


	var (
		db *gorm.DB = config.SetupDatabaseConnection()

		userRepository repository.UserRepository = repository.NewUserRepository(db)

		userService service.UserService = service.NewUserService(userRepository)

		userController controllers.UserController = controllers.NewUserController(userService)
	)

	router := gin.Default()

	routes.UserRoutes(router, userController)

	router.Run()
}