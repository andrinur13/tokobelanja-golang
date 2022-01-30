package main

import (
	"fmt"
	"tokobelanja-golang/conf"
	"tokobelanja-golang/controller"
	"tokobelanja-golang/middleware"
	"tokobelanja-golang/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"tokobelanja-golang/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db := conf.InitDB()

	// repository
	userRepository := repository.NewUserRepository(db)

	// service
	userService := service.NewUserService(userRepository)

	// controller
	userController := controller.NewUserController(userService)

	router := gin.Default()

	// routing
	// user
	router.POST("/register", userController.RegisterUser)
	router.POST("/login", userController.Login)
	router.PUT("/update-account", middleware.AuthMiddleware(), userController.UpdateUser)
	router.DELETE("/delete-account", middleware.AuthMiddleware(), userController.DeleteUser)

	router.Run()

}
