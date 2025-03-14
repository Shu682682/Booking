package handlers

import (
	"net/http"

	"github.com/Shu682682/Booking.git/pkg/handlers/config"
	"github.com/Shu682682/Booking.git/pkg/handlers/render"
	"github.com/Shu682682/Booking.git/pkg/models"
)

//Repo the repository used by the handlers
var Repo *Repository
//Repository is the repository type
type Repository struct{
	App *config.AppConfig

}
//New Repor creats a new repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App:a,
	}
}

//NewHandlers sets the repoitory for the handlers
func NewHandlers(r *Repository){
	Repo=r
}

//Home is the home page handler
func (m *Repository)Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.html", &models.TemplateData{})
	
}
//About is the about page handler
func(m *Repository) About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about.html", &models.TemplateData{} )
}

func(m*Repository)Index(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "index.html", &models.TemplateData{})

}
func(m*Repository)General(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "generals_quarter.html", &models.TemplateData{})

}
func(m*Repository)Major(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "majors_suite.html", &models.TemplateData{})

}

func(m*Repository)Contact(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "contact.html", &models.TemplateData{})

}
func(m*Repository)Book(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "book.html", &models.TemplateData{})

}
	



