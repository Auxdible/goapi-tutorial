package handlers

import (
	"github.com/auxdible/goapi-tutorial/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
