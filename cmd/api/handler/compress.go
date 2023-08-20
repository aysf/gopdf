package handler

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func Compress(c echo.Context) error {

	w, _ := os.Getwd()
	filePath := "/storage/pdf/yle-flyers-word-list-picture-book-2018.pdf"
	filePath2 := "/storage/pdf/yle-flyers-word-list-picture-book-2018_compressed.pdf"

	fullPath := w + filePath
	fullPath2 := w + filePath2

	conf := api.LoadConfiguration()
	conf.OptimizeDuplicateContentStreams = true
	conf.WriteXRefStream = true

	api.OptimizeFile(fullPath, fullPath2, conf)

	fs, _ := os.Open(fullPath)
	defer fs.Close()

	ctx, err := api.ReadContextFile(fullPath2)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = api.OptimizeContext(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  "success",
	})
}
