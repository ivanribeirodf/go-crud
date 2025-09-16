package main

import (
	"log"
    "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
    "go-crud/database"
    "go-crud/models"
    "go-crud/routes"
	"go-crud/middlewares"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Println("Arquivo .env não encontrado, usando variáveis do sistema")
    }

    r := gin.Default()

	r.Use(middlewares.ValidationErrorHandler)

    database.ConnectDB()
    database.DB.AutoMigrate(&models.User{})

    routes.SetupRoutes(r)

    r.Run(":8080")
}