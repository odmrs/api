package main

import (
	"fmt"
	"net/http"
	"os/exec" // To clear the console

	// Framework CHI
	//My handlers files
	// Lib to use logs
	// Logs are registers of events inside our api

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus" // The 'log' will be the alias to use this lib

	"github.com/odmrs/api/internal/handlers"
)

func main() {
	log.SetReportCaller(true)        // Turn on
	var r *chi.Mux = chi.NewRouter() // Pointer to this struct used in the api
	handlers.Handler(r)
	exec.Command("clear")
	fmt.Println("Runing GO API")

	// Create a server
	err := http.ListenAndServe("localhost:8000", r) // The port to server run and the handler
	if err != nil {
		log.Error(err)
	}
}
