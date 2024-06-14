package controllers

import (
	"nebula_backend/config"
	"nebula_backend/models"
	"nebula_backend/serializers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Users retrieved successfully",
		"data":    serializers.SerializeUsers(users),
	})
}
