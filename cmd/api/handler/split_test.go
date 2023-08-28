package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestPdfSplit(t *testing.T) {

	e := echo.New()

	jsonStr := `{
		"name": "camry_ebrochure.pdf",
		"path": "/storage/testPdf",
		"selected_pages": ["1-3","5","8-9"]
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/pdf/split", strings.NewReader(jsonStr))
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := Split(c); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	t.Cleanup(func() {
		w, _ := os.Getwd()
		os.Remove(w + "/storage/testPdf" + "/camry_ebrochure_split_1-3.pdf")
		os.Remove(w + "/storage/testPdf" + "/camry_ebrochure_split_5.pdf")
		os.Remove(w + "/storage/testPdf" + "/camry_ebrochure_split_8-9.pdf")
	})

}
