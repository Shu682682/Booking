package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler{
	csrfHandler :=nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:"/",
		Secure:false,
		SameSite:http.SameSiteLaxMode,
	})
	return csrfHandler
}
//SessionLoad loads and saves the seesion on every request
func SessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}