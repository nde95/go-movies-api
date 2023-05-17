package main

import (
	"encoding/json"
	"fmt"
	"moviesapi/internal/models"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Cinema is up and running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "2014-10-26")

	interstellar := models.Movie{
		ID:          1,
		Title:       "Interstellar",
		ReleaseDate: rd,
		MPAARating:  "PG-13",
		RunTime:     169,
		Description: "Interstellar movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, interstellar)

	rd, _ = time.Parse("2006-01-02", "1995-06-30")

	apollo13 := models.Movie{
		ID:          2,
		Title:       "Apollo 13",
		ReleaseDate: rd,
		MPAARating:  "PG",
		RunTime:     140,
		Description: "Apollo 13 movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, apollo13)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
