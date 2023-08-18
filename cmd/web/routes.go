package main

import (
	"net/http"

	"github.com/bagasmad/bookings/pkg/config"
	"github.com/bagasmad/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// using pat
// func routes(app *config.AppConfig) http.Handler {
// 	//multiplexer is called mux
// 	//package really good at routing and very simple one
// 	mux := pat.New()
// 	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
// 	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
// 	return mux
// }

//package really good at routing and very simple one

// using chi
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	//use recoverer middleware to absorb panic
	mux.Use(middleware.Recoverer)
	//use middleware that we just made
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	//go get static files from
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
