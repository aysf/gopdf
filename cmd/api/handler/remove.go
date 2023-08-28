package handler

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type RemoveData struct {
	InName      string   `json:"inname" validate:"required"`
	InPath      string   `json:"inpath" validate:"required"`
	OutName     string   `json:"outname"`
	OutPath     string   `json:"outpath"`
	TargetPages []string `json:"target_pages"`
}

func Remove(c echo.Context) error {

	rd := new(RemoveData)

	if err := c.Bind(rd); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	w, _ := os.Getwd()
	inFile := w + rd.InPath + "/" + rd.InName

	var outFile string
	if rd.OutPath != "" && rd.OutName != "" {
		outFile = w + rd.OutPath + "/" + rd.OutName
	}

	err := api.RemovePagesFile(inFile, outFile, rd.TargetPages, api.LoadConfiguration())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  "success",
	})

}
