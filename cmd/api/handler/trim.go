package handler

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type (
	TrimData struct {
		InName      string   `json:"inname" validate:"required"`
		InPath      string   `json:"inpath" validate:"required"`
		OutName     string   `json:"outname"`
		OutPath     string   `json:"outpath"`
		TargetPages []string `json:"target_pages"`
	}
)

func Trim(c echo.Context) error {

	td := new(TrimData)
	if err := c.Bind(td); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	w, _ := os.Getwd()
	inFile := w + td.InPath + "/" + td.InName

	var outFile string
	if td.OutPath != "" && td.OutName != "" {
		outFile = w + td.OutPath + "/" + td.OutName
	}

	err := api.TrimFile(inFile, outFile, td.TargetPages, api.LoadConfiguration())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  "success",
	})
}
