package handler

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type (
	CompressData struct {
		InFile  string `json:"infile"`
		InPath  string `json:"inpath"`
		OutFile string `json:"outfile"`
		OutPath string `json:"outpath"`
	}
)

func Compress(c echo.Context) error {

	cd := new(CompressData)
	if err := c.Bind(cd); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	w, _ := os.Getwd()
	inFile := w + cd.InPath + "/" + cd.InFile
	outFile := w + cd.OutPath + "/" + cd.OutFile

	conf := api.LoadConfiguration()
	conf.OptimizeDuplicateContentStreams = true
	conf.WriteXRefStream = true

	err := api.OptimizeFile(inFile, outFile, conf)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// fs, _ := os.Open(inFile)
	// defer fs.Close()

	// ctx, err := api.ReadContextFile(outFile)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	// err = api.OptimizeContext(ctx)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  "success",
	})
}
