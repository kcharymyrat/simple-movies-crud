package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func getAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bson, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bson)
}

func getMovieByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	for _, movie := range movies {
		fmt.Printf("movie.ID = %s, id = %s\n", movie.ID.String(), id)
		if movie.ID.String() == id {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	errMsg := map[string]string{"error": "Movie with id =" + id + "does not exits"}
	json.NewEncoder(w).Encode(errMsg)
}

func createMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")

	var newMovie Movie

	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newMovie.ID = uuid.New()
	movies = append(movies, newMovie)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMovie)
}

func updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// read the id from request
	id := chi.URLParam(r, "id")
	idUUID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Not correct ID = "+id, http.StatusBadRequest)
		return
	}

	// Check whether movie with this id exists - if so delete
	deleted := false
	var result []Movie
	for _, movie := range movies {
		if movie.ID.String() != id {
			result = append(result, movie)
		}
		if movie.ID.String() == id {
			deleted = true
		}
	}

	// If no such movie
	if !deleted {
		http.Error(w, "No such movie with ID ="+id, http.StatusBadRequest)
		return
	}

	// Get the body of request
	var updatedMovie Movie
	err = json.NewDecoder(r.Body).Decode(&updatedMovie)
	if err != nil {
		http.Error(w, "Wrong movie json content", http.StatusBadRequest)
		return
	}

	updatedMovie.ID = idUUID

	movies = result
	movies = append(movies, updatedMovie)

	bson, err := json.Marshal(updatedMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bson)
}

func deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// read the id from request
	id := chi.URLParam(r, "id")

	// Check whether movie with this id exists - if so delete
	deleted := false
	var result []Movie
	for _, movie := range movies {
		if movie.ID.String() != id {
			result = append(result, movie)
		}
		if movie.ID.String() == id {
			deleted = true
		}
	}

	// If no such movie
	if !deleted {
		http.Error(w, "No such movie with ID = "+id, http.StatusBadRequest)
		return
	}

	movies = result

	w.WriteHeader(http.StatusOK)
}
