package routes

import "net/http"

// Route for a custom api call
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
