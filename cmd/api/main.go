package main

import (
	"fmt"
	"log"
	campaign "marketplace/internal/campaign/domain"
	"marketplace/internal/campaign/infrastructure/db"
	infra_mail "marketplace/internal/campaign/infrastructure/mail"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

type Application struct {
	service campaign.ServiceInterface
}

func (app *Application) server() error {
	srv := &http.Server{
		Addr: ":8000",
		Handler: app.routes(),//routes configurado em route.go
		IdleTimeout: 30 * time.Second,
		ReadTimeout: 10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	fmt.Println("Server running on port 8000...")

	return srv.ListenAndServe()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env")
	}
	service := campaign.Service{
		Repository: &db.Repository{
			DB: db.NewDB(),
		},
		SendMail: infra_mail.SendMail,
	}
	
	app := Application{
		service: &service,
	}

	err = app.server()
	if err != nil {
		log.Fatal(err)
	}
}