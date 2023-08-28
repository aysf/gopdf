package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestReorder(t *testing.T) {
	defer cleanup(t)

	e := echo.New()

	jsonStr := `{
		"inname": "yle-flyers-sample.pdf",
		"inpath": "/storage/testPdf",
		"outname": "yle-flyers-sample_reordered.pdf",
		"outpath": "/storage/testPdf",
		"new_page_order": ["3-31","1-2","32-36"]
	}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/pdf/split", strings.NewReader(jsonStr))
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := Reorder(c); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

}

func cleanup(t *testing.T) {
	t.Cleanup(func() {
		w, _ := os.Getwd()
		os.Remove(w + "/storage/testPdf" + "/yle-flyers-sample_reordered.pdf")
		os.Remove(w + "/storage/testPdf" + "/temp_yle-flyers-sample.pdf")
		os.RemoveAll(w + "/storage/testPdf/yle-flyers-sample_dir")
	})
}
