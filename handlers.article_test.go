package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	// Create a request to send to the above route
	r.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := io.ReadAll(w.Body)
		pageOK := err == nil && strings.Contains(string(p), "<title>Home Page</title>")

		return statusOK && pageOK
	})
}
