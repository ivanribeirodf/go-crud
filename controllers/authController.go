package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"go-crud/database"
	"go-crud/dto"
	"go-crud/models"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

func Register(c *gin.Context) {
	var input dto.CreateUserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Passwd), bcrypt.DefaultCost)
	role := "user"
	if input.Role != "" {
		role = input.Role
	}
	user := models.User{
		Name:   input.Name,
		Email:  input.Email,
		Passwd: string(hashedPassword),
		Role:   role,
	}
	database.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário registrado com sucesso"})
}

func Login(c *gin.Context) {
	var input dto.LoginDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou senha inválidos"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(input.Passwd)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou senha inválidos"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(jwtKey)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
