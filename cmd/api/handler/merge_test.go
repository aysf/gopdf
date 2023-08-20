package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestPdfMerge(t *testing.T) {

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/pdf/merge", strings.NewReader(jsonStr))
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

}
