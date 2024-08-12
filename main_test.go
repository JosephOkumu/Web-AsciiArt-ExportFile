package main

import (
	"asciiart/web"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(web.HomeHandler)

	// Serve the HTTP request to the ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect (200 OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestAsciiArtHandler(t *testing.T) {
	// Create a request with POST method and form values
	formData := strings.NewReader("text=hello&banner=thinkertoy")
	req, err := http.NewRequest("POST", "/ascii-art", formData)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(web.AsciiArtHandler)

	// Serve the HTTP request to the ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect (200 OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
