package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestJpgToPdf(t *testing.T) {

	e := echo.New()

	jsonStr := `{
		"infiles": [
		  {
			"name": "dhiva-krishna-YApS6TjKJ9c-unsplash.jpg",
			"path": "/storage/testImage"
		  },
		  {
			"name": "dima-panyukov-DwxlhTvC16Q-unsplash.jpg",
			"path": "/storage/testImage"
		  },
		  {
			"name": "kenny-eliason-FcyipqujfGg-unsplash.jpg",
			"path": "/storage/testImage"
		  }
		],
		"outfile": "jpg-to-pdf-output.pdf",
		"outpath": "/storage/testImage",
		"configs": {
		  "page_size": "A4",
		  "scale": 0.95
		}
	  }`

	req := httptest.NewRequest(http.MethodPost, "/api/v1/pdf/jpg-to-pdf", strings.NewReader(jsonStr))
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := JpgToPdf(c); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	t.Cleanup(func() {
		w, _ := os.Getwd()
		os.Remove(w + "/storage/testImage" + "/jpg-to-pdf-output.pdf")
	})
}
