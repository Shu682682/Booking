package render

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Shu682682/Booking.git/internal/config"
	"github.com/Shu682682/Booking.git/internal/models"
	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}
var app *config.AppConfig
var pathToTemplates="./templates"
var session *scs.SessionManager


// NewTemplates initializes app configuration
func NewTemplates(a *config.AppConfig) {
	app = a

	var err error
	app.TemplateCache, err = CreateTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache:", err)
	}
}

func AddDefaultData(td *models.TemplateData,r *http.Request) *models.TemplateData{
	td.CSRFToken = nosurf.Token(r)
	

	if td == nil {
		td = &models.TemplateData{}
	}
	if td.StringMap == nil {
		td.StringMap = make(map[string]string)
	}

	if session != nil {
		td.Flash = session.PopString(r.Context(), "flash")
	}

	return td
}

// RenderTemplate renders an HTML template with a base template
func RenderTemplate(w http.ResponseWriter, r *http.Request, html string, td *models.TemplateData) error{
	var tc map[string]*template.Template
	var err error

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Println("Error creating new template cache:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			
		}
	}

	// Get requested template from cache
	t, ok := tc[html]
	if !ok {
		log.Println("Could not get template from template cache:", html)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return errors.New("Cant get template from cache")
	}

	buf := new(bytes.Buffer)
	td =AddDefaultData(td,r)

	err = t.Execute(buf,td)


	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Fatal(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to response:", err)
	}
	return nil
}

// CreateTemplateCache generates a map of templates for caching
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.html", pathToTemplates)) 
	if err != nil {
		log.Println("Error finding page templates:", err)
		return myCache, err
	}
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

