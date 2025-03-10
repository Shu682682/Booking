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

	remoteIP:=r.RemoteAddr//every time someone hit the home page will keep the IP
	m.App.Session.Put(r.Context(),"remote_ip",remoteIP)

	render.RenderTemplate(w, "home.html", &models.TemplateData{})
	
}
//About is the about page handler
func(m *Repository) About(w http.ResponseWriter, r *http.Request){
	// //perform some logic 
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	td := &models.TemplateData{
		StringMap: stringMap,
	}

	remoteIP:=m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"]=remoteIP
	render.RenderTemplate(w, "about.html", td)
}

	



