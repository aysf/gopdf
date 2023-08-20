package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestPdfSplit(t *testing.T) {

	m, _ := os.Getwd()
	log.Println("cek wd1: ", m)

	e := echo.New()

	jsonStr := `{"name":"camry_ebrochure.pdf","path":"/storage/pdf", "range":"1-3,7-9"}`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/pdf/split", strings.NewReader(jsonStr))
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := PdfSplit(c); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

}
