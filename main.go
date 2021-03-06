package main

import (
	"github.com/edwinnduti/go-postgres/router"
	"log"
	"os"
	"net/http"
)

func main() {
	// call router
	r := router.Router()

	// establish port number
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8030"
	}

	// set server
	server := &http.Server{
		Handler: r, // n
		Addr:    ":" + Port,
	}

	// log server output
	log.Printf("Listening on PORT: %s", Port)
	server.ListenAndServe()
}
