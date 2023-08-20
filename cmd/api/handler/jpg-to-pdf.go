package handler

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

const (
	dirPath        = "/storage/image"
	dirPathTest    = "/storage/testImage"
	imageFileTest1 = "dhiva-krishna-YApS6TjKJ9c-unsplash.jpg"
	imageFileTest2 = "dima-panyukov-DwxlhTvC16Q-unsplash.jpg"
	imageFileTest3 = "kenny-eliason-FcyipqujfGg-unsplash.jpg"
)

func JpgToPdf(c echo.Context) error {

	w, _ := os.Getwd()
	filePath := w + dirPath
	if1, if2, if3 := filePath+"/"+imageFileTest1, filePath+"/"+imageFileTest2, filePath+"/"+imageFileTest3

	var ifs []string

	ifs = append(ifs, if1, if2, if3)

	imp := pdfcpu.Import{
		PageDim:  types.PaperSize["A4"],
		PageSize: "A4",
		Pos:      types.Center,
		Scale:    0.95,
		InpUnit:  types.POINTS,
	}

	err := api.ImportImagesFile(ifs, filePath+"/jpg-to-pdf-output.pdf", &imp, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  "success",
	})
}
