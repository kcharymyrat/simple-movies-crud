package main

import (
	"github.com/google/uuid"
)

type Movie struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Year     int       `json:"year"`
	Director *Director `json:"director"`
}

type Director struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}
