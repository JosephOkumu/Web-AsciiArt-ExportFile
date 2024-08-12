package web

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

type Ascii struct {
	Result string
}

// asciiArtHandler handles requests to the "/ascii-art" path
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request to %s\n", r.URL.Path)

	if r.URL.Path != "/ascii-art" {
		renderError(w, http.StatusNotFound, "", "Not Found", "The page you are looking for could not be found.")
		return
	}

	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		if strings.Contains(text, "\r\n") {
			text = strings.ReplaceAll(text, "\r\n", "\n")
		}
		banner := r.FormValue("banner")

		// Generate ASCII art using the provided text and banner
		result, err := generateAsciiArt(text, banner)
		if err != nil {
			if strings.Contains(err.Error(), "tampered with") {
				renderError(w, http.StatusInternalServerError, err.Error(), "Internal Server Error", "The banner file: "+banner+" is corrupted")
				return
			} else if strings.Contains(err.Error(), "invalid input") {
				renderError(w, http.StatusInternalServerError, err.Error(), "Internal Server Error", "The input text contains non-ASCII characters that could not be converted to ASCII art")
				return
			}
			renderError(w, http.StatusNotFound, err.Error(), "Not Found", "The banner you are looking for could not be found.")
			return
		}

		t, err := template.ParseFiles("templates/result.html")
		if err != nil {
			renderError(w, http.StatusNotFound, err.Error(), "Not Found", "The template you are looking for could not be found.")
			return
		}

		ascii := Ascii{
			Result: result,
		}

		t.Execute(w, ascii)
	} else {
		renderError(w, http.StatusBadRequest, "", "Bad Request", "Your request could not be processed.")
	}
}
