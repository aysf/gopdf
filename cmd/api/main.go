package main

import (
	"net/http"

	"github.com/aysf/gopdf/cmd/api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.Route(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
