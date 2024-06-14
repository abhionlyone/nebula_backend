package controllers

import (
	"nebula_backend/config"
	"nebula_backend/models"
	"nebula_backend/serializers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	last24Hours := time.Now().Add(-24 * time.Hour)
	if err := config.DB.Where("created_at >= ?", last24Hours).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userCount := len(users)
	c.JSON(http.StatusOK, gin.H{
		"count": userCount,
		"users": serializers.SerializeUsers(users),
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, serializers.SerializeUser(user))
}
