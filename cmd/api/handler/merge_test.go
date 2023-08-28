package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestPdfMerge(t *testing.T) {

	e := echo.New()

	jsonStr := `{"infiles":[{"name":"camry_ebrochure.pdf","path":"/storage/pdf"},{"name":"mirai_ebrochure.pdf","path":"/storage/pdf"}], "outfile":"camry_mirai_ebrochure.pdf"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/pdf/merge", strings.NewReader(jsonStr))
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := Merge(c); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	t.Cleanup(func() {
		w, _ := os.Getwd()
		os.Remove(w + "/storage/pdf" + "/camry_mirai_ebrochure.pdf")
	})

}
