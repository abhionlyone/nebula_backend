package routes

import (
	"nebula_backend/controllers"
	"nebula_backend/customMiddleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	apiV1 := e.Group("/api/v1")
	apiV1.GET("/users", controllers.GetUsers, customMiddleware.LogRequest)
	apiV1.GET("/users/:id", controllers.GetUser)
	return e
}
