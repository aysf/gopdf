package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type (
	MergeData struct {
		InFiles []struct {
			FileName string `json:"name" validate:"required"`
			FilePath string `json:"path" validate:"required"`
		} `json:"infiles"`
		OutFile string `json:"outfile"`
	}
)

func Merge(c echo.Context) error {

	md := new(MergeData)

	temp := c.Request().Body
	dec := json.NewDecoder(temp)
	dec.Decode(md)

	w, _ := os.Getwd()
	inFiles := []string{}

	for i := 0; i < len(md.InFiles); i++ {
		inFiles = append(inFiles, w+md.InFiles[i].FilePath+"/"+md.InFiles[i].FileName)
	}

	var outFile string = md.OutFile
	if outFile == "" {
		outFile = w + md.InFiles[0].FilePath + "/" + "merged_file"
	} else {
		t := w + md.InFiles[0].FilePath + "/" + outFile
		outFile = t
	}

	err := api.MergeCreateFile(inFiles, outFile, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  "success",
	})
}

// func mergePDF() error {

// 	return nil
// }
