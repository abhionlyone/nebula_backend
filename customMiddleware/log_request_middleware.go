package customMiddleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

func LogRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Printf("Handling request: %s %s", c.Request().Method, c.Request().RequestURI)
		return next(c)
	}
}
