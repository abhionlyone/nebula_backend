package controllers

import (
	"nebula_backend/config"
	"nebula_backend/models"
	"nebula_backend/serializers"
	"net/http"
	"strconv"

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

func GetUser(c echo.Context) error {
	// Get the user ID from the URL parameter
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User retrieved successfully",
		"data":    serializers.SerializeUser(user),
	})
}
