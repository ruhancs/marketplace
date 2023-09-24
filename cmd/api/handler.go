package main

import (
	"encoding/json"
	"fmt"
	"marketplace/internal/contract"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (app *Application) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var req contract.CreateCampaignDTO
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		return
	}
	//pegar o email do contexto retirado do jwt e informa que o contexto é uma string
	email := r.Context().Value("email").(string)
	//inserir o email em req
	req.CreatedBy = email

	id, err := app.service.Create(req)
	if err != nil {
		render.Status(r,400)
		render.JSON(w,r,map[string]string{"err": err.Error()})
		return
	}

	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"id": id})
}

func (app *Application) GetAllCampaign(w http.ResponseWriter, r *http.Request) {
	campaigns := app.service.GetAll()

	render.Status(r, 200)
	render.JSON(w, r, campaigns)
}

func (app *Application) GetCampaignByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	campaign,err := app.service.GetOne(id)
	if err != nil {
		fmt.Println(err)
		render.Status(r, 404)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, campaign)
}

func (app *Application) CancelCampaignByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := app.service.Cancel(id)
	if err != nil {
		fmt.Println(err)
		render.Status(r, 404)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, map[string]string{"status": "campaign canceled"})
}

func (app *Application) DeleteCampaign(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := app.service.Delete(id)
	if err != nil {
		fmt.Println(err)
		render.Status(r, 404)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, map[string]string{"status": "campaign deleted"})
}
