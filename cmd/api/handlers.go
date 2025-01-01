package api

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type Response struct {
	URL string
}

func (api *API) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func (api *API) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/response.html"))
	originalURL := r.FormValue("url")
	if !isValidURL(originalURL) {
		http.Error(w, "invalid URL", http.StatusBadRequest)
		return
	}

	shortCode, err := api.GenerateUniqueShortCode()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	_, err = api.DB.Connection.Exec(
		"INSERT INTO urls (short_code, original_url) VALUES ($1, $2)",
		shortCode,
		originalURL,
	)
	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	shortURL := r.Host + "/" + shortCode

	rsp := Response{URL: shortURL}
	tmpl.Execute(w, rsp)
}

func (api *API) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	if shortCode == "" {
		http.Error(w, "broken url", http.StatusBadRequest)
	}

	var originalURL string
	err := api.DB.Connection.QueryRow("SELECT original_url FROM urls WHERE short_code = $1;", shortCode).
		Scan(&originalURL)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
