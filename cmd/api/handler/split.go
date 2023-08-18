package handler

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func PdfSplit(c echo.Context) error {

	sd := new(SplitData)
	if err := c.Bind(sd); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	w, _ := os.Getwd()
	filePath := w + sd.Path + "/"
	fileName := sd.Name

	f, err := os.Open(filePath + fileName)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"error": "error opening file: " + err.Error(),
			"data":  nil,
		})
	}

	rn, err := splitRange(sd.Range)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "error split range: " + err.Error(),
			"data":  nil,
		})
	}

	f2, err := addBookmark(f, filePath, fileName, rn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "error adding bookmark: " + err.Error(),
			"data":  nil,
		})
	}
	defer f2.Close()

	outDir := filePath

	outputFileName := fileNameWithoutExt(sd.Name) + "_split"

	err = api.Split(f2, outDir, outputFileName, 0, api.LoadConfiguration())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "error split pdf: " + err.Error(),
			"data":  nil,
		})
	}

	removeFileTemp(filePath)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  "success",
	})

}

func splitRange(input string) ([][]int, error) {
	groups := strings.Split(input, ",")
	result := make([][]int, len(groups))

	for i, group := range groups {
		ranges := strings.Split(group, "-")
		start, _ := strconv.Atoi(ranges[0])

		if len(ranges) == 1 {
			result[i] = []int{start}
		} else {
			end, _ := strconv.Atoi(ranges[1])
			if start > end {
				return nil, errors.New("invalid range")
			}
			result[i] = []int{start, end}
		}
	}

	return result, nil
}

func addBookmark(f *os.File, filePath string, fileName string, r [][]int) (*os.File, error) {

	outputPath := filePath + "temp_" + fileName
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating output PDF:", err)
		return nil, err
	}

	b := []pdfcpu.Bookmark{}

	for k, v := range r {
		fn := fileName + "_split_" + strconv.Itoa(v[0])
		if len(v) == 2 {
			fn += "-" + strconv.Itoa(v[1])
		}

		b = append(b, pdfcpu.Bookmark{
			PageFrom: v[0],
			Title:    fn,
		})
		if len(v) == 1 {
			b = append(b, pdfcpu.Bookmark{
				PageFrom: v[0] + 1,
				Title:    "temp_" + strconv.Itoa(k),
			})
		} else {
			b = append(b, pdfcpu.Bookmark{
				PageFrom: v[1] + 1,
				Title:    "temp_" + strconv.Itoa(k),
			})
		}

	}

	// fmt.Println(b)

	api.AddBookmarks(f, outputFile, b, false, api.LoadConfiguration())

	return outputFile, f.Close()
}

func fileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func removeFileTemp(dir string) error {

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasPrefix(info.Name(), "temp_") {
			err := os.Remove(path)
			if err != nil {
				return err
			}
			fmt.Println("Removed:", path)
		}

		return nil

	})

	return err
}
