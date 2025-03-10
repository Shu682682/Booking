package main

//after add the handler package type go run cmd/web/*.go or go run ./cmd/web
import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Shu682682/Booking.git/pkg/handlers"
	"github.com/Shu682682/Booking.git/pkg/handlers/config"
	"github.com/Shu682682/Booking.git/pkg/handlers/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080" // cannot change use const
var app config.AppConfig
var session *scs.SessionManager

// main is the main Application function
func main() {
	
	//change this to true when in production
	app.InProduction=false

	session=scs.New()
	session.Lifetime=24* time.Hour
	session.Cookie.Persist=true
	session.Cookie.SameSite=http.SameSiteLaxMode
	session.Cookie.Secure=app.InProduction
	app.Session=session

	// 建立 template cache
	tc, err := render.CreateTemplateCache() 
	if err != nil {
		log.Fatal("Cannot create template cache:", err) //
	}
	app.TemplateCache = tc
	app.UseCache=false

	repo:=handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//initial template
	render.NewTemplates(&app) //


	// initial server
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber)) 

	srv :=&http.Server{
		Addr: portNumber,
		Handler: routes(&app),
		
	}
	err =srv.ListenAndServe()
	log.Fatal(err)

}
