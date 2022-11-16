package server

import (
	"groupie-tracker-vizualization/cmd/logic"
	"groupie-tracker-vizualization/cmd/server/errors"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errors.Errors(w, r, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		errors.Errors(w, r, http.StatusMethodNotAllowed)
		return
	}

	Data, wrong := logic.AllArtists(w, r)
	if wrong {
		errors.Errors(w, r, http.StatusInternalServerError)
		return
	}

	files := []string{"ui/templates/base.html", "ui/templates/home.html", "ui/templates/block.html"}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		errors.Errors(w, r, http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", Data)
	if err != nil {
		errors.Errors(w, r, http.StatusInternalServerError)
		return
	}
}

func detailArtist(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		errors.Errors(w, r, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		errors.Errors(w, r, http.StatusMethodNotAllowed)
		return
	}

	Data, wrong := logic.DetailArtist(w, r)
	if wrong != 0 {
		if wrong == 404 {
			errors.Errors(w, r, http.StatusNotFound)
		} else if wrong == 500 {
			errors.Errors(w, r, http.StatusInternalServerError)
		} else {
			errors.Errors(w, r, http.StatusInternalServerError)
		}
		return
	}

	files := []string{"ui/templates/base.html", "ui/templates/details.html"}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		errors.Errors(w, r, http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", Data)
	if err != nil {
		errors.Errors(w, r, http.StatusInternalServerError)
		return
	}
}
