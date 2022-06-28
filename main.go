package main

import (
	"fmt"
	"goOwl/routers"
	"goOwl/utilities"

	"github.com/labstack/echo/v4"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	fmt.Printf("%s | %s\n", c.Path(), err)
}

func main() {

	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Renderer = utilities.T

	e.Static("/public", "public")
	e.File("/sw.js", "public/sw.js")
	e.File("/robots.txt", "public/robots.txt")
	e.File("/manifest.webmanifest", "public/manifest.webmanifest")

	routers.BlogRouter(e)
	routers.SeriesRouter(e)
	routers.CategoryRouter(e)
	routers.AboutRouter(e)

	e.Logger.Fatal(e.Start(":3000"))
}
