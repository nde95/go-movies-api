package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// application config
	var app application

	// read from CLI

	// connect to postgres

	app.Domain = "example.com"

	log.Println("starting server on port", port)

	// start webserver
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
