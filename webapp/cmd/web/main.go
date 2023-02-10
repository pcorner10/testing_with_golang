package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {
	// set up and app config
	app := application{}

	// get aplications routes
	mux := app.routes()

	// print out a message
	log.Panicln("Starting server on port 8080...")

	// start the server
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
