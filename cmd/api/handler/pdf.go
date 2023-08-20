package handler

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func PdfHandler(c echo.Context) error {

	path := c.QueryParam("path")
	name := c.QueryParam("name")

	w, _ := os.Getwd()

	mc, err := api.ReadContextFile(w + path + "/" + name)
	if err != nil {
		return c.String(http.StatusBadRequest, "error: "+string(err.Error()))

	}

	data := map[string]interface{}{
		"author":     mc.Author,
		"creator":    mc.Creator,
		"name":       mc.Names,
		"created_at": mc.CreationDate,
		"pages":      mc.PageCount,
		"title":      mc.Title,
		"producer":   mc.Producer,
		"mod_date":   mc.ModDate,
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  data,
	}, "    ")
}
