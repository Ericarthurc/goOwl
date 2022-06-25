package main

import (
	"fmt"
	"goOwl/routers"
	"goOwl/utilities"
	"net/http"

	"github.com/labstack/echo/v4"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
}

func main() {

	e := echo.New()
	e.HideBanner = true
	// e.HTTPErrorHandler = customHTTPErrorHandler
	e.Renderer = utilities.T

	e.Static("/public", "public")
	e.File("/sw.js", "public/sw.js")
	e.File("/robots.txt", "public/robots.txt")
	e.File("/manifest.webmanifest", "public/manifest.webmanifest")

	routers.BlogRouter(e)
	routers.SeriesRouter(e)
	routers.CategoryRouter(e)
	routers.AboutRouter(e)

	e.Logger.Fatal(e.Start(":1323"))
}
