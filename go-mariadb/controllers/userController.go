package controllers

import (
	"github.com/Animesh-roy100/go-mariadb/db"
	"github.com/Animesh-roy100/go-mariadb/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	db.DBUser.Find(&users)
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.DBUser.Create(&user)
	c.JSON(200, user)
}

func GetUser(c *gin.Context) {
	var user models.User
	if err := db.DBUser.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := db.DBUser.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.DBUser.Save(&user)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := db.DBUser.Where("id = ?", c.Param("id")).Delete(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, gin.H{"result": "User deleted"})
}
