package routes

import (
	"github.com/aysf/gopdf/cmd/api/handler"
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	api := e.Group("/api")
	v1 := api.Group("/v1")
	pdf := v1.Group("/pdf")

	pdf.GET("/info", handler.Info)
	pdf.POST("/split", handler.Split)
	pdf.POST("/merge", handler.Merge)
	pdf.POST("/jpg-to-pdf", handler.JpgToPdf)
	pdf.POST("/compress", handler.Compress)
	pdf.POST("/trim", handler.Trim)
	pdf.POST("/remove", handler.Remove)
	pdf.POST("/reorder", handler.Reorder)

}
