package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

func LogControllerAction(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Log the request URI and method
		log.Printf("Handling request: %s %s", c.Request().Method, c.Request().RequestURI)

		// Call the next handler
		return next(c)
	}
}
