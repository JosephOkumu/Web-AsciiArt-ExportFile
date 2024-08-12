package web

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"time"
)

// DownloadHandler handles requests to the "/download" path
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request to %s\n", r.URL.Path)

	if r.URL.Path != "/download" {
		renderError(w, http.StatusNotFound, "", "Not Found", "The page you are looking for could not be found.")
		return
	}

	if r.Method == http.MethodGet {
		format := r.URL.Query().Get("format")
		content := r.URL.Query().Get("content")

		var filename string

		if content == "" {
			renderError(w, http.StatusBadRequest, "", "Bad Request", "No content to download.")
			return
		}

		switch format {
		case "txt":
			filename = generateFilename(format)
			w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Length", strconv.Itoa(len(content)))

			// Set Cache-Control header
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")

			w.Write([]byte(content))
		case "html":
			filename = generateFilename(format)
			content = convertToHTML(content)
			w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
			w.Header().Set("Content-Type", "text/html")
			w.Header().Set("Content-Length", strconv.Itoa(len(content)))

			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")

			w.Write([]byte(content))
		default:
			renderError(w, http.StatusBadRequest, "", "Bad Request", "Unsupported file format.")
			return
		}
	} else {
		renderError(w, http.StatusMethodNotAllowed, "", "Method Not Allowed", "Invalid request method.")
	}
}

// convertToHTML converts ASCII art to HTML format
func convertToHTML(asciiArt string) string {
	return "<!DOCTYPE html>" +
		"<html lang=\"en\">" +
		"<head><meta charset=\"UTF-8\"><title>ASCII Art</title></head>" +
		"<body><pre>" + html.EscapeString(asciiArt) + "</pre></body>" +
		"</html>"
}

// generateFilename creates a timestamped filename based on the requested format
func generateFilename(format string) string {
	timestamp := time.Now().Format("20060102150405") // Format: YYYYMMDDHHMMSS
	return fmt.Sprintf("ascii_art_%s.%s", timestamp, format)
}
