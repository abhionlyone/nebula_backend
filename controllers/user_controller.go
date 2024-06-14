package controllers

import (
	"fmt"
	"nebula_backend/config"
	"nebula_backend/models"
	"nebula_backend/serializers"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if len(users) < 10 {
		timestamp := time.Now().Format("20060102150405")
		newUser := models.User{
			Name:  "Random User",
			Email: fmt.Sprintf("randomuser%s@example.com", timestamp),
		}
		if err := config.DB.Create(&newUser).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		users = append(users, newUser)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Users retrieved successfully",
		"data":    serializers.SerializeUsers(users),
	})
}
