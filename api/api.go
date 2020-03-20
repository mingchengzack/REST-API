package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// API has router
type API struct {
	Router     *mux.Router
	Subrouters map[string]*mux.Router
}

// Init the api with routers
func (a *API) Init() {
	a.Router = mux.NewRouter()
	a.Subrouters = make(map[string]*mux.Router)
}

// Run the api on its router
func (a *API) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
