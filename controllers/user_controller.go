package controllers

import (
	"nebula_backend/config"
	"nebula_backend/models"
	"nebula_backend/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, serializers.SerializeUsers(users))
}
