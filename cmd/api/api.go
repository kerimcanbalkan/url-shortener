package api

import (
	"github.com/gorilla/mux"

	"github.com/kerimcanbalkan/url-shortener/db"
)

type API struct {
	DB     *db.DB
	Router *mux.Router
}

func NewAPI(database *db.DB) *API {
	router := mux.NewRouter()
	api := &API{DB: database, Router: router}
	api.setupRoutes()

	return api
}

func (api *API) setupRoutes() {
	api.Router.HandleFunc("/", api.IndexHandler).Methods("GET")
	api.Router.HandleFunc("/shorten", api.ShortenHandler).Methods("POST")
	api.Router.HandleFunc("/{shortCode}", api.RedirectHandler).Methods("GET")
}
