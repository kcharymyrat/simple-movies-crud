package main

import (
	"net/http"

	"github.com/google/uuid"
)

var movies []Movie

func main() {
	d1 := Director{Name: "Scorsese", Country: "USA"}
	d2 := Director{Name: "Guy Ritchie", Country: "UK"}

	movies = []Movie{
		{ID: uuid.New(), Title: "Godfather", Year: 1988, Director: &d1},
		{ID: uuid.New(), Title: "Revolver", Year: 2003, Director: &d2},
	}

	// Ensure the `routes` function is defined
	mux := routes()

	// Setup a server
	srv := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	// Listen and Serve
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
