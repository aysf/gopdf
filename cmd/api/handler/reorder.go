package handler

import (
	"errors"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type (
	ReorderData struct {
		InName       string   `json:"inname"`
		InPath       string   `json:"inpath"`
		OutName      string   `json:"outname"`
		OutPath      string   `json:"outpath"`
		NewPageOrder []string `json:"new_page_order"`
	}
)

func Reorder(c echo.Context) error {

	var err error
	rd := new(ReorderData)
	if err := c.Bind(rd); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("error bind parameters: "+err.Error()))

	}

	w, _ := os.Getwd()
	inPath := w + rd.InPath + "/"
	inName := rd.InName
	outDir := inPath + fileNameWithoutExt(rd.InName)
	splitDir := outDir + "_dir"

	_ = createDirIfNotExist(splitDir)

	// check input
	err = SplitPDF(rd.NewPageOrder, inPath, inName, splitDir)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("error splitting pdf: "+err.Error()))
	}
	removeFileTemp(outDir)
	removeFileTemp(splitDir)

	// merge the trimmed file with pdf that page already removed
	var inFiles []string

	err = filepath.Walk(splitDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			inFiles = append(inFiles, splitDir+"/"+info.Name())
		}
		return nil
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("error collecting splitted files: "+err.Error()))
	}

	err = api.MergeCreateFile(inFiles, w+rd.OutPath+"/"+rd.OutName, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("error merging splitted files: "+err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  "success",
	})

}
