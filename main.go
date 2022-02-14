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
	transactionRepository := repository.NewTransactionHistoryRepository(db)
	productRepository := repository.NewProductRepository(db)

	// service
	userService := service.NewUserService(userRepository)
	transactionService := service.NewTransactionHistoryService(transactionRepository, productRepository, userRepository)

	// controller
	userController := controller.NewUserController(userService)
	transactionController := controller.NewTransactionHistoryController(transactionService, userService)

	router := gin.Default()

	// routing
	// user
	router.POST("users/register", userController.RegisterUser)
	router.POST("users/login", userController.Login)
	router.POST("users/topup", middleware.AuthMiddleware(), userController.TopUpSaldo)

	// transaction
	router.POST("transactions", middleware.AuthMiddleware(), transactionController.NewTransaction)
	router.POST("transactions/my-transactions", middleware.AuthMiddleware(), transactionController.GetMyTransaction)

	router.Run()

}
