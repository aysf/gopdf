package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCompress(t *testing.T) {

	e := echo.New()

	jsonStr := `
	{
		"infile": "camry_ebrochure.pdf",
		"inpath": "/storage/testPdf",
		"outfile": "camry_ebrochure_compressed.pdf",
		"outpath": "/storage/testPdf"
	}`

	req := httptest.NewRequest(http.MethodPost,
		"/api/v1/pdf/compress",
		strings.NewReader(jsonStr))
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := Compress(c); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	t.Cleanup(func() {
		w, _ := os.Getwd()
		os.Remove(w + "/storage/testPdf" + "/camry_ebrochure_compressed.pdf")
	})

}
