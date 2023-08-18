package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// we will insert the handler
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		//serve the http writer and reader
		next.ServeHTTP(w, r)
	})
}

// Add CSRF protection to all POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   appConfig.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// Loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
