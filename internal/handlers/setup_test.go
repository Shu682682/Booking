package handlers

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/Shu682682/Booking.git/internal/config"
	"github.com/Shu682682/Booking.git/internal/models"
	"github.com/Shu682682/Booking.git/internal/render"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justinas/nosurf"
)


var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates="./../templates"
var functions = template.FuncMap{}
	


func init() {
    cwd, _ := os.Getwd()
    pathToTemplates = filepath.Join(cwd, "../../templates") //calculate correct path
    log.Println("Updated pathToTemplates:", pathToTemplates)
}

func getRoutes() http.Handler{
	gob.Register(models.Reservation{})
	app.InProduction=false

	session=scs.New()
	session.Lifetime=24* time.Hour
	session.Cookie.Persist=true
	session.Cookie.SameSite=http.SameSiteLaxMode
	session.Cookie.Secure=app.InProduction
	app.Session=session

	// 建立 template cache
	tc, err :=CreateTestTemplateCache() 
	if err != nil {
		log.Fatal("Cannot create template cache:", err) 
		
	}
	app.TemplateCache = tc
	app.UseCache=true

	repo:=NewRepo(&app)
	NewHandlers(repo)

	//initial template
	render.NewTemplates(&app) 

	mux :=chi.NewRouter()
	mux.Use(middleware.Recoverer)
	//mux.Use(NoSurf)//write basic middleware 
	mux.Use(SessionLoad)
	mux.Get("/",Repo.Home)
	mux.Get("/about",Repo.About)
	mux.Get("/index",Repo.Index)
	mux.Get("/rooms/generals-quarter",Repo.General)
	mux.Get("/rooms/majors-suite",Repo.Major)
	mux.Get("/contact",Repo.Contact)

	mux.Get("/book",Repo.Book)
	mux.Post("/book",Repo.PostBook)
	mux.Post("/book-json",Repo.AvailabilityJSON)

	// mux.Get("/book",Repo.Reservation)
	// mux.Post("/book",Repo.PostReservation)



	fileServer:=http.FileServer(http.Dir("./static/"))//connect to image
	mux.Handle("/static/*", http.StripPrefix("/static",fileServer))

	return mux

}

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




func CreateTestTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	wd, _ := os.Getwd()
    log.Println("Current working directory:", wd)
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.html", pathToTemplates)) 
	if err != nil {
		log.Println("Error finding page templates:", err)
		return myCache, err
	}

	
	log.Println("Templates found:", pages)  // 确保找到了模板
	// log.Println("Found page templates:", pages)

	for _, page := range pages {
		name := filepath.Base(page)
		// log.Println("Parsing template:", name)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			log.Println("Error parsing template:", name, err)
			continue
		}

		// 找 base layout (base.layout.html)
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.html",pathToTemplates))
		if err != nil {
			log.Println("Error finding layout templates:", err)
			continue
		}

		if len(matches) > 0 {
			// log.Println("Parsing layout templates:", matches)
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
			if err != nil {
				log.Println("Error parsing layout templates:", err)
				continue
			}
		}

		if ts != nil { 
			myCache[name] = ts
		}
	}
	return myCache, nil
}
