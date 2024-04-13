package handlers

import (
	// middlewares help the comunication and the integration between diferents part of a system

	"github.com/go-chi/chi"
	chimmidle "github.com/go-chi/chi/middleware"

	"github.com/odmrs/api/internal/middleware"
)

func Handler(r *chi.Mux) {
  r.Use(chimmidle.StripSlashes) // This will ignore all StripSlashes

  r.Route("/account", func(router chi.Router) {
    // Middleware for /account route 
    router.Use(middleware.Authorization) // Will authorize this router

    // Get the coin balance

    router.Get("/coins", GetCoinBalance)
  })
}


