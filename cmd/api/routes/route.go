package routes

import (
	"github.com/aysf/gopdf/cmd/api/handler"
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	api := e.Group("/api")
	v1 := api.Group("/v1")
	pdf := v1.Group("/pdf")

	pdf.GET("/info", handler.PdfHandler)
	pdf.POST("/split", handler.PdfSplit)

}
