package handler

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type (
	ImgData struct {
		InFiles []struct {
			FileName string `json:"name" validate:"required"`
			FilePath string `json:"path" validate:"required"`
		} `json:"infiles"`
		OutFile string `json:"outfile"`
		OutPath string `json:"outpath"`
		Configs struct {
			PageSize string  `json:"page_size"`
			Scale    float64 `json:"scale"`
		} `json:"configs"`
	}
)

func JpgToPdf(c echo.Context) error {

	id := new(ImgData)
	if err := c.Bind(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	w, _ := os.Getwd()
	of := w + id.OutPath + "/" + id.OutFile

	var ifs []string

	for i := 0; i < len(id.InFiles); i++ {
		ifs = append(ifs, w+id.InFiles[i].FilePath+"/"+id.InFiles[i].FileName)
	}

	imp := pdfcpu.Import{
		PageDim:  types.PaperSize["A4"],
		PageSize: "A4",
		Pos:      types.Center,
		Scale:    0.95,
		InpUnit:  types.POINTS,
	}

	if id.Configs.PageSize != "" {
		imp.PageSize = id.Configs.PageSize
		imp.PageDim = types.PaperSize[id.Configs.PageSize]
	}
	if id.Configs.Scale != 0.0 {
		imp.Scale = id.Configs.Scale
	}

	err := api.ImportImagesFile(ifs, of, &imp, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  "success",
	})
}
