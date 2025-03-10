package main

import (
	"net/http"

	"github.com/Shu682682/Booking.git/pkg/handlers"
	"github.com/Shu682682/Booking.git/pkg/handlers/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler{
	mux :=chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)//write basic middleware 
	mux.Use(SessionLoad)
	mux.Get("/",handlers.Repo.Home)
	mux.Get("/about",handlers.Repo.About)

	return mux
}

