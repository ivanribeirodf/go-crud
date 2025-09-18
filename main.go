package main

import (
	"go-crud/database"
	"go-crud/middlewares"
	"go-crud/models"
	"go-crud/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go-crud/docs"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis do sistema")
	}

	r := gin.Default()

	r.Use(middlewares.ValidationErrorHandler)
	r.Use(middlewares.LoggerMiddleware())

	database.ConnectDB()
	database.DB.AutoMigrate(&models.User{})

	routes.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(r.Run(":8080"))
}
