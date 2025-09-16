package routes

import (
    "github.com/gin-gonic/gin"
    "go-crud/controllers"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/users", controllers.CreateUser)
    r.GET("/users", controllers.GetUsers)
    r.GET("/users/:id", controllers.GetUser)
    r.PUT("/users/:id", controllers.UpdateUser)
    r.DELETE("/users/:id", controllers.DeleteUser)
}