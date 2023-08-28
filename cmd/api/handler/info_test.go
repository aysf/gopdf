package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestInfo(t *testing.T) {

	e := echo.New()

	jsonStr := `
	{
		"infile": "camry_ebrochure.pdf",
		"inpath": "/storage/testPdf",
		"outfile": "camry_ebrochure_compressed.pdf",
		"outpath": "/storage/testPdf"
	}`

	req := httptest.NewRequest(http.MethodGet,
		"/api/v1/pdf/compress",
		strings.NewReader(jsonStr))

	q := req.URL.Query()

	q.Add("path", "/storage/testPdf")
	q.Add("name", "camry_ebrochure.pdf")

	req.URL.RawQuery = q.Encode()

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := Info(c); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

}
