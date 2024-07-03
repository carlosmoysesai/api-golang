package controller

import (
	"github.com/carlosmoysesai/api-golang/config"
	"github.com/carlosmoysesai/api-golang/models"
	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context) {
	c.JSON(200,
		gin.H{
			"message": "bem vindo ao controle de user!",
		})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&user)
	c.JSON(201, user)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.Where("id =?", id).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&user)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "Usuario não existe"})
		return
	}
	c.JSON(200, gin.H{"message": "Usuario deletado"})
}
