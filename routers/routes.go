package routers

import (
	"fmt"
	"goOwl/middlewares"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func BlogRouter(e *echo.Echo) {
	blogs := e.Group("/")
	blogs.GET("", func(c echo.Context) error {
		dt := time.Now()
		fmt.Println("Handler | " + dt.Format("01-02-2006 15:04:05"))
		return c.String(http.StatusOK, "blogIndex")
	}, middlewares.TestMiddle)
	blogs.GET("blog", func(c echo.Context) error {
		return c.String(http.StatusOK, "blogPost")
	})
}

func SeriesRouter(e *echo.Echo) {
	series := e.Group("/series")
	series.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "series")
	})
}

func CategoryRouter(e *echo.Echo) {
	category := e.Group("/category")
	category.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "category")
	})
}

func AboutRouter(e *echo.Echo) {
	about := e.Group("/about")
	about.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "about")
	})
}
