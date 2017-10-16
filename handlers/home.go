package handlers

import (
	"html/template"
	"net/http"

	"github.com/derekslenk/s_k-web/libhttp"
)

type stats struct {
	Episodes  int
	AvgLength int
}

// GetHome handles the hompage
func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/home.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}
	stats := stats{
		Episodes:  68,
		AvgLength: 55,
	}
	tmpl.Execute(w, stats)
}
