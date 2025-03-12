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
	mux.Get("/index",handlers.Repo.Index)
	mux.Get("/rooms/generals-quarter",handlers.Repo.General)
	mux.Get("/rooms/majors-suite",handlers.Repo.Major)
	mux.Get("/contact",handlers.Repo.Contact)
	mux.Get("/book",handlers.Repo.Book)


	fileServer:=http.FileServer(http.Dir("./static/"))//connect to image
	mux.Handle("/static/*", http.StripPrefix("/static",fileServer))

	return mux
}

