package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"github.com/derekslenk/s_k-web/libhttp"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type stats struct {
	EpisodeCount        int `json:"EpisodeCount"`
	SpecialEpisodeCount int `json:"SpecialEpisodeCount"`
	AvgLength           int `json:"AvgLength"`
	AvgLengthSpecial    int `json:"AvgLengthSpecial"`
}

// GetHome handles the hompage
func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/home.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	s, err := queryStats()

	tmpl.Execute(w, s)
}

func queryStats() (stats, error) {
	resp, err := myClient.Get("http://steveandkyle.accountant/api/stats")
	if err != nil {
		return stats{}, err
	}

	defer resp.Body.Close()

	var d stats

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return stats{}, err
	}

	return d, nil
}
