package web

import (
	"fmt"
	"net/http"
	"text/template"
)

// homeHandler handles requests to the root path ("/")
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request to %s\n", r.URL.Path)

	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound, "", "Not Found", "The page you are looking for could not be found.")
		return
	}

	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			renderError(w, http.StatusNotFound, err.Error(), "Not Found", "The template you are looking for could not be found.")
			return
		}

		t.Execute(w, nil)
	} else {
		renderError(w, http.StatusBadRequest, "", "Bad Request", "Your request could not be processed.")
	}
}
