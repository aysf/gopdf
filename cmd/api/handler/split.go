package handler

import (
	"errors"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type (
	SplitData struct {
		Name          string   `json:"name" validate:"required"`
		Path          string   `json:"path" validate:"required"`
		Range         string   `json:"range"`
		SelectedPages []string `json:"selected_pages"`
	}
)

var filePath string
var fileName string

func Split(c echo.Context) error {

	var err error

	sd := new(SplitData)
	if err := c.Bind(sd); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	w, _ := os.Getwd()
	filePath = w + sd.Path + "/"
	fileName = sd.Name
	outDir := filePath + fileNameWithoutExt(fileName) + "_dir"

	_ = createDirIfNotExist(outDir)

	err = SplitPDF(sd.SelectedPages, filePath, fileName, outDir)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "error spliting pages: " + err.Error(),
			"data":  nil,
		})
	}

	removeFileTemp(filePath)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  "success",
	})

}

func SplitPDF(pages []string, filePath, fileName, outDir string) error {
	var err error

	f, err := os.Open(filePath + fileName)

	if err != nil {
		return errors.New("error opening file: " + err.Error())
	}

	rn, om, err := strToIntArr(pages)

	if err != nil {
		return errors.New("error converting page(s) string to array: " + err.Error())
	}
	_, err = flatten2D(rn)
	if err != nil {
		return errors.New("range not valid: " + err.Error())
	}

	mc, err := api.ReadContextFile(filePath + fileName)
	if err != nil {
		return errors.New("error reading context: " + err.Error())
	}

	f2, err := addBookmark(f, filePath, fileName, rn, om, mc)
	if err != nil {
		return errors.New("error bookmarking page(s): " + err.Error())
	}
	defer f2.Close()

	if outDir == "" {
		outDir = filePath + "/" + fileNameWithoutExt(fileName) + "_dir"
	}

	// err = api.Split(f2, outDir, outputFileName, 0, api.LoadConfiguration())
	err = api.SplitFile(filePath+"temp_"+fileName, outDir, 0, api.LoadConfiguration())
	if err != nil {
		return errors.New("error split API: " + err.Error())
	}
	return nil
}

func strToIntArr(groups []string) ([][]int, map[int]int, error) {

	result := make([][]int, len(groups))
	originalOrderMap := make(map[int]int)

	for i, group := range groups {
		ranges := strings.Split(group, "-")
		start, _ := strconv.Atoi(ranges[0])

		if len(ranges) == 1 {
			result[i] = []int{start}
		} else {
			end, _ := strconv.Atoi(ranges[1])
			if start > end {
				return nil, originalOrderMap, errors.New("invalid range")
			}
			result[i] = []int{start, end}
		}
	}

	// Create a map to store the original order
	for idx, value := range result {
		originalOrderMap[value[0]] = idx
	}

	// order the 2d int array
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result, originalOrderMap, nil
}

func flatten2D(slice [][]int) ([]int, error) {
	var result []int

	for _, innerSlice := range slice {
		if len(innerSlice) == 2 {
			start, end := innerSlice[0], innerSlice[1]

			if start > end {
				return nil, errors.New("invalid input: start value is greater than end value")
			}

			for i := start; i <= end; i++ {
				result = append(result, i)
			}
		} else if len(innerSlice) == 1 {
			result = append(result, innerSlice[0])
		} else {
			return nil, errors.New("invalid input: inner slice must have 1 or 2 elements")
		}
	}

	// Check for duplicate elements
	seen := make(map[int]bool)
	for _, num := range result {
		if seen[num] {
			return nil, errors.New("invalid input: non-unique element in the result slice")
		}
		seen[num] = true
	}

	return result, nil
}

func addBookmark(f *os.File, filePath string, fileName string, groups [][]int, orderMap map[int]int, fileCtx *model.Context) (*os.File, error) {

	outputPath := filePath + "temp_" + fileName
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return nil, err
	}

	b := []pdfcpu.Bookmark{}

	totalPage := fileCtx.PageCount

	for i := 0; i < len(groups); i++ {
		curGroup := groups[i]
		// nextGroup := groups[i+1]

		fn := zeroPadding(orderMap[curGroup[0]], 3) + "_-" + fileNameWithoutExt(fileName) + "_split_" + strconv.Itoa(curGroup[0])
		if len(curGroup) == 2 {
			fn += "-" + strconv.Itoa(curGroup[1])
		}

		b = append(b, pdfcpu.Bookmark{
			PageFrom: curGroup[0],
			Title:    fn,
		})

		if curGroup[len(curGroup)-1] == totalPage {
			break
		}

		b = append(b, pdfcpu.Bookmark{
			PageFrom: curGroup[len(curGroup)-1] + 1,
			Title:    "temp_" + strconv.Itoa(i),
		})

	}

	err = api.AddBookmarks(f, outputFile, b, false, api.LoadConfiguration())
	if err != nil {
		return nil, err
	}

	return outputFile, f.Close()
}

func zeroPadding(input int, totalDigits int) string {
	inputStr := strconv.Itoa(input)

	paddingCount := totalDigits - len(inputStr)

	paddedStr := ""
	for i := 0; i < paddingCount; i++ {
		paddedStr += "0"
	}

	paddedStr += inputStr

	return paddedStr
}

func fileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func createDirIfNotExist(path string) error {

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(path, os.ModePerm)
		// if err != nil {
		// 	return fmt.Errorf("failed creating directory: %v", err)
		// }
	}

	return nil
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
		}

		return nil

	})

	return err
}
