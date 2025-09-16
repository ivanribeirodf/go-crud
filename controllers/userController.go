package controllers

import (
	"github.com/gin-gonic/gin"
    "go-crud/database"
    "go-crud/models"
    "net/http"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
    var users []models.User
    database.DB.Find(&users)
    c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
    var user models.User
    id := c.Param("id")
    result := database.DB.First(&user, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
    var user models.User
    id := c.Param("id")

    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
        return
    }

    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Model(&user).Updates(input)
    c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
    var user models.User
    id := c.Param("id")
    if err := database.DB.Delete(&user, id).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao deletar"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado"})
}