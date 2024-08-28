package main

import (
	"golang-extract-api/src/configuration/database"
	"golang-extract-api/src/controllers"
	"golang-extract-api/src/repositories"
	"golang-extract-api/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	connection := database.NewConnection()

	extractRepository := repositories.NewExtractRepository(connection.DB)
	extractService := services.NewExtractService(extractRepository)

	clientRepository := repositories.NewClientRepository(connection.DB)
	clientService := services.NewClientService(clientRepository, extractService)

	clientController := controllers.NewClientController(clientService, extractService)

	server := gin.Default()

	server.GET("/clients/:id/extrato", clientController.GetExtract)
	server.POST("/clients/:id/transacoes", clientController.UpdateCreditUsed)

	server.Run(":9999")
}
