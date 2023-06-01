package main

import (
	"Web_Service/internal/delivery/api/handlers"
	"Web_Service/internal/infrastructure/service"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	randomUserService := service.NewRandomUserService()

	apiHandler := handlers.NewAPIHandler(randomUserService)

	router.GET("/api/person", apiHandler.GetPerson)

	router.Run(":8080")
}
