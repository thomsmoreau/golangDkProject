package handlers

import (
	"golangdk/views"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func FrontPage(mux chi.Router) {
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// We ignore the error returned from Render for now, because there's currently nothing sensible we could do with it.
		_ = views.FrontPage().Render(w)
	})
}
