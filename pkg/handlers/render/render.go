package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Shu682682/Booking.git/pkg/handlers/config"
	"github.com/Shu682682/Booking.git/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates initializes app configuration
func NewTemplates(a *config.AppConfig) {
	app = a

	var err error
	app.TemplateCache, err = CreateTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache:", err)
	}
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	if td == nil {
		td = &models.TemplateData{}
	}
	if td.StringMap == nil {
		td.StringMap = make(map[string]string) 
	}
	return td
}

// RenderTemplate renders an HTML template with a base template
func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {
	var tc map[string]*template.Template
	var err error

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Println("Error creating new template cache:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	// Get requested template from cache
	t, ok := tc[html]
	if !ok {
		log.Println("Could not get template from template cache:", html)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)
	td =AddDefaultData(td)

	err = t.Execute(buf,td)


	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to response:", err)
	}
}

// CreateTemplateCache generates a map of templates for caching
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	
	pages, err := filepath.Glob("./templates/*.html") 
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
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			log.Println("Error finding layout templates:", err)
			continue
		}

		if len(matches) > 0 {
			// log.Println("Parsing layout templates:", matches)
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				log.Println("Error parsing layout templates:", err)
				continue
			}
		}

		if ts != nil { 
			myCache[name] = ts
			log.Println("Cached template:", name)
		}
	}
	return myCache, nil
}

