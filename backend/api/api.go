package api

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/maryam-nokohan/WeatherApp/backend/config"
)

type Application struct {
	Config config.Config
}

func (app *Application) Mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Use(middleware.Timeout(60 * time.Second))
	
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthHandler)

	})

	return r
}

func (app *Application) Run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server is running...at %s", app.Config.Addr)

	return srv.ListenAndServe()
}
