package main

import (
	"ki_assignment-1/config"
	"ki_assignment-1/controllers"
	"ki_assignment-1/repository"
	"ki_assignment-1/routes"
	"ki_assignment-1/service"

	"github.com/gin-contrib/cors"
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

		jwtService service.JWTService = service.NewJWTService()

		userRepository repository.UserRepository = repository.NewUserRepository(db)
		fileRepository repository.FileRepository = repository.NewFileRepository(db)

		userService service.UserService = service.NewUserService(userRepository)
		fileService service.FileService = service.NewFileService(fileRepository, userRepository)

		userController controllers.UserController = controllers.NewUserController(userService, jwtService)
		fileController controllers.FileController = controllers.NewFileController(fileService)
	)

	router := gin.Default()

	router.Use(cors.Default())

	routes.UserRoutes(router, userController, jwtService)
	routes.FileRoutes(router, fileController, jwtService)

	router.Run()
}
