package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hadirafad/bookings/pkg/config"
	"github.com/hadirafad/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	filServer := http.FileServer(http.Dir("D:/golang codes/course 2/booking/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", filServer))

	return mux
}
