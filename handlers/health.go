package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Health(mux chi.Router) {
	/*
		Pass mux directly to facilitate tests, route closer to function
		URL's path often contains variables, nice to be able to see the name close to the code
	*/
	// Default to returning HTTP 200 OK if nothing else is set
	mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {

	})
}
