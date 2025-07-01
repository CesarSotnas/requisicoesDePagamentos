package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/CesarSotnas/requisicoesDePagamentos.git/handlers"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Post("/payment", handlers.ProcessPayment)
	r.Get("/stats", handlers.GetStats)

	return r
}
