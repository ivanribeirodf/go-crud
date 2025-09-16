package controllers

import (
    "net/http"
    "go-crud/database"
    "go-crud/models"
    "go-crud/dto"

    "github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
    var input dto.CreateUserDTO
    if err := c.ShouldBindJSON(&input); err != nil {
        // c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Error(err)
        return
    }

    user := models.User{Name: input.Name, Email: input.Email}
    database.DB.Create(&user)

    c.JSON(http.StatusCreated, gin.H{"data": user})
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
    var input dto.UpdateUserDTO
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
        return
    }

    database.DB.Model(&user).Updates(models.User{Name: input.Name, Email: input.Email})
    c.JSON(http.StatusOK, gin.H{"data": user})
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
