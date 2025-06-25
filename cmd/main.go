package main

import (
	"nexa/internal/database"
	"nexa/internal/handler"
	"nexa/internal/service"
	"nexa/internal/utils"
)

func main() {
	utils.LoadEnv()
	pool := database.ConnectDB()
	userHandler := handler.NewUserHandler(pool)
	transactionHandler := handler.NewTransactionHandler(pool)
	categoryHandler := handler.NewCategoryHandler(pool)
	service.StartServer(userHandler, transactionHandler, categoryHandler)
}
