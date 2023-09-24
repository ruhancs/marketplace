package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	//proteger rotas com auth
	mux.Route("/campaign", func(r chi.Router) {
		mux.Use(app.Auth)
		mux.Post("/", app.CreateCampaign)
		mux.Patch("/cancel/{id}", app.CancelCampaignByID)
		mux.Delete("/{id}", app.DeleteCampaign)
	})

	mux.Get("/campaign", app.GetAllCampaign)
	mux.Get("/campaign/{id}", app.GetCampaignByID)

	return mux
}
