package main

import (
	"encoding/json"
	"fmt"
	"marketplace/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (app *Application) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var req contract.CreateCampaignDTO
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		return
	}

	id, err := app.service.Create(req)
	if err != nil {
		render.Status(r,400)
		render.JSON(w,r,map[string]string{"err": err.Error()})
		return
	}

	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"id": id})
}
