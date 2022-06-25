package middlewares

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func TestMiddle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		dt := time.Now()
		fmt.Println("Test Middleware | " + c.Path() + " | " + dt.Format("01-02-2006 15:04:05"))
		return next(c)
	}
}
