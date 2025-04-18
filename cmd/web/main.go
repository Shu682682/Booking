package main

//after add the handler package type go run cmd/web/*.go or go run ./cmd/web
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Shu682682/Booking.git/internal/config"
	"github.com/Shu682682/Booking.git/internal/handlers"
	"github.com/Shu682682/Booking.git/internal/helpers"
	"github.com/Shu682682/Booking.git/internal/render"
	"github.com/alexedwards/scs/v2"
)


const portNumber = ":8085" 
var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger


// main is the main Application function
func main() {
	err:=run()
	if err!=nil{
		log.Fatal(err)
	}
	
	//change this to true when in production
	
	// initial server
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber)) 

	srv :=&http.Server{
		Addr: portNumber,
		Handler: routes(&app),
		
	}
	err =srv.ListenAndServe()
	log.Fatal(err)

}
func run()error{
	app.InProduction=false

	infoLog=log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog=infoLog

	errorLog=log.New(os.Stdout,"ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog=errorLog

	session=scs.New()
	session.Lifetime=24* time.Hour
	session.Cookie.Persist=true
	session.Cookie.SameSite=http.SameSiteLaxMode
	session.Cookie.Secure=app.InProduction
	app.Session=session

	// 建立 template cache
	tc, err := render.CreateTemplateCache() 
	if err != nil {
		log.Fatal("Cannot create template cache:", err) 
		return err
	}
	app.TemplateCache = tc
	app.UseCache=false

	repo:=handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	//initial template
	render.NewTemplates(&app) 
	helpers.NewHelpers(&app)


	return nil
}
