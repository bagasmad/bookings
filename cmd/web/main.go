package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bagasmad/bookings/pkg/config"
	"github.com/bagasmad/bookings/pkg/handlers"
	"github.com/bagasmad/bookings/pkg/render"
)

// since middleware is in the same package as main
var appConfig config.AppConfig

// use it for middleware
var session *scs.SessionManager

func main() {

	//change this to true when it's in production
	appConfig.InProduction = true

	session = scs.New()
	//set the time in session
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	appConfig.Session = session

	//set pointer

	templateCacheCreated, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	appConfig.TemplateCache = templateCacheCreated
	appConfig.UseCache = true
	render.AccessConfig(&appConfig)

	repo := handlers.NewRepo(&appConfig)
	handlers.SetRepo(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// _ = http.ListenAndServe(":8080", nil)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&appConfig),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
