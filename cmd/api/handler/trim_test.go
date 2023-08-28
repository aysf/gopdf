package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestTrim(t *testing.T) {

	e := echo.New()

	jsonStr := `{
		"inname": "yle-flyers-sample.pdf",
		"inpath": "/storage/testPdf",
		"outname": "yle-flyers-sample_removed.pdf",
		"outpath": "/storage/testPdf",
		"target_pages": ["1-3","5","8-9"]
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/pdf/remove", strings.NewReader(jsonStr))
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := Trim(c); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	t.Cleanup(func() {
		w, _ := os.Getwd()
		os.Remove(w + "/storage/testImage" + "/yle-flyers-sample_trimmed.pdf")
	})
}
