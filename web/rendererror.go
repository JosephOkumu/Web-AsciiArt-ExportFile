package web

import (
	"net/http"
	"text/template"
)

// renderError renders an error response with the specified status code and message
func renderError(w http.ResponseWriter, status int, errorMessage string, message string, information string) {
	w.WriteHeader(status)

	var templateFile string = "templates/status_codes.html"

	t, err := template.ParseFiles(templateFile)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Status  int
		Message string
		Error   string
		Info    string
	}{
		Status:  status,
		Message: message,
		Error:   errorMessage,
		Info:    information,
	}

	t.Execute(w, data)
}
