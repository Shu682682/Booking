package render

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Shu682682/Booking.git/internal/models"
	"github.com/alexedwards/scs/v2"
)




func TestAddDefaultData(t *testing.T) {
	if session == nil {
		session = scs.New()
		testApp.Session = session
	}

	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Fatal("Failed to create session request:", err)
	}

	session.Put(r.Context(), "flash", "123")

	flash := session.GetString(r.Context(), "flash")
	if flash != "123" {
		t.Fatalf("Session value not stored properly, expected '123' but got: '%s'", flash)
	}

	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Errorf("flash not found in session, expected '123' but got: '%s'", result.Flash)
	}
}

func TestRenderTemplate(t *testing.T){
	pathToTemplates = "./../../templates"
	// pathToTemplates="./../templates"
	tc, err:=CreateTemplateCache()
	if err!=nil{
		t.Error(err)
	}
	app.TemplateCache=tc
	r,err:=getSession()
	if err!=nil{
		t.Error(err)
	}
	rr:=httptest.NewRecorder()

	err=RenderTemplate(rr,r, "index.html",&models.TemplateData{})
	if err!=nil{
		t.Errorf("error writing template to browser: %v", err)
	}
	err = RenderTemplate(rr, r, "non-existent.html", &models.TemplateData{})
	if err == nil {
		t.Error("expected error for non-existent template, but got none")
	}
	for k := range tc {
		t.Log("Template loaded:", k)
	}

	// var ww myWriter
	// err=RenderTemplate(&ww,r,"index.html",&models.TemplateData{})
	// if err!=nil{
	// 	t.Error("error writing template to browser")
	// }
	// err=RenderTemplate(&ww,r,"non-existent.html",&models.TemplateData{})
	// if err!=nil{
	// 	t.Error("rendered template that does not exist")
	// }
}


func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx, err := session.Load(r.Context(), "")
	if err != nil {
		return nil, err
	}

	r = r.WithContext(ctx) 
	return r, nil
}



func TestNewTemplates(t *testing.T){
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T){
	pathToTemplates="./../../templates"
	_,err:=CreateTemplateCache()
	if err!=nil{
		t.Error(err)
	}
}