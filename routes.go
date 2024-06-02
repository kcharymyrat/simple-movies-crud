package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/movies", getAllMoviesHandler)
	r.Get("/movies/{id}", getMovieByIDHandler)
	r.Post("/movies", createMovieHandler)
	r.Put("/movies/{id}", updateMovieHandler)
	r.Delete("/movies/{id}", deleteMovieHandler)

	return r
}
