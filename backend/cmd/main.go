package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/maryam-nokohan/WeatherApp/backend/api"
	"github.com/maryam-nokohan/WeatherApp/backend/config"
	"github.com/maryam-nokohan/WeatherApp/backend/internal/handlers"
)

func main() {
	// load env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}
	config.Load()

	app := &api.Application{
		Config: config.Config{
			Addr: config.AppConfig.Addr,
		},
	}

	r := chi.NewRouter()

	r.Mount("/v1" , app.Mount())

	// Routes
	r.Get("/" , handlers.HomePage)
	r.Post("/search" , handlers.SearchCityPage)
	r.Get("/city" , handlers.HandleCity)

fs := http.FileServer(http.Dir("frontend"))
    r.Handle("/static/*", http.StripPrefix("/static/", fs))

    if err := app.Run(r); err != nil {
        log.Fatal(err)
    }
}