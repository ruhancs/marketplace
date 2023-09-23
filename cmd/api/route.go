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

	mux.Post("/campaign", app.CreateCampaign)
	mux.Get("/campaign", app.GetAllCampaign)
	mux.Get("/campaign/{id}", app.GetCampaignByID)
	mux.Patch("/campaign/cancel/{id}", app.CancelCampaignByID)
	mux.Delete("/campaign/{id}", app.DeleteCampaign)

	return mux
}
