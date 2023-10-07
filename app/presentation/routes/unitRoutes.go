package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func unitRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", getUnits)
	return r
}

func getUnits(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Route 2"))
}
