package handlers

import (
	"html/template"
	"net/http"

	"github.com/derekslenk/s_k-web/libhttp"
)

func GetEpisodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/episodes.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	tmpl.Execute(w, nil)
}
