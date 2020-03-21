package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// API has router
type API struct {
	Router     *mux.Router
	Subrouters map[string]*mux.Router
}

// NewAPI creates a new api
func NewAPI() *API {
	api := &API{}
	api.Router = mux.NewRouter()
	api.Subrouters = make(map[string]*mux.Router)
	return api
}

// AddSubRouter adds a subrouter with given prefix
func (api *API) AddSubRouter(path string) {
	api.Subrouters[path] = api.Router.PathPrefix(path).Subrouter()
}

// SetSubRoute sets a route a subrouter
func (api *API) SetSubRoute(subroute, path, method string, f http.HandlerFunc) {
	if _, ok := api.Subrouters[subroute]; !ok {
		api.AddSubRouter(subroute)
	}
	api.Subrouters[subroute].HandleFunc(path, f).Methods(method)
}

// SetRoute sets a specific route
func (api *API) SetRoute(path, method string, f http.HandlerFunc) {
	api.Router.HandleFunc(path, f).Methods(method)
}

// SubGet wraps the api for subrouter GET method
func (api *API) SubGet(subroute, path string, f http.HandlerFunc) {
	if _, ok := api.Subrouters[subroute]; !ok {
		api.AddSubRouter(subroute)
	}
	api.Subrouters[subroute].HandleFunc(path, f).Methods("GET")
}

// SubPost wraps the api for subrouter POST method
func (api *API) SubPost(subroute, path string, f http.HandlerFunc) {
	if _, ok := api.Subrouters[subroute]; !ok {
		api.AddSubRouter(subroute)
	}
	api.Subrouters[subroute].HandleFunc(path, f).Methods("POST")
}

// SubPut wraps the api for subrouter PUT method
func (api *API) SubPut(subroute, path string, f http.HandlerFunc) {
	if _, ok := api.Subrouters[subroute]; !ok {
		api.AddSubRouter(subroute)
	}
	api.Subrouters[subroute].HandleFunc(path, f).Methods("PUT")
}

// SubDelete wraps the api for subrouter DELETE method
func (api *API) SubDelete(subroute, path string, f http.HandlerFunc) {
	if _, ok := api.Subrouters[subroute]; !ok {
		api.AddSubRouter(subroute)
	}
	api.Subrouters[subroute].HandleFunc(path, f).Methods("DELETE")
}

// Get wraps the api for GET method
func (api *API) Get(path string, f http.HandlerFunc) {
	api.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the api for POST method
func (api *API) Post(path string, f http.HandlerFunc) {
	api.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the api for PUT method
func (api *API) Put(path string, f http.HandlerFunc) {
	api.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the api for DELETE method
func (api *API) Delete(path string, f http.HandlerFunc) {
	api.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the api on its router
func (api *API) Run(host string) {
	fmt.Println("Starting server on the port" + host + "...")
	log.Fatal(http.ListenAndServe(host, api.Router))
}
