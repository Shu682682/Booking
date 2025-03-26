package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Shu682682/Booking.git/internal/config"
	"github.com/Shu682682/Booking.git/internal/models"
	"github.com/alexedwards/scs/v2"
)

var testApp config.AppConfig

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})
	testApp.Session = session


	session = scs.New()//initial session
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.InProduction = false
	testApp.Session = session

	tc, err := CreateTemplateCache() 
	if err != nil {
		log.Fatal("Cannot create template cache:", err)
	}

	testApp.TemplateCache = tc
	testApp.UseCache = false

	NewTemplates(&testApp)

	exitCode := m.Run()

	os.Exit(exitCode)
}

type myWriter struct{

}

func(tw *myWriter) Header() http.Header{
	var h http.Header
	 return h
}

func (tw *myWriter) WriteHeader(i int ){

}

func (tw *myWriter) Write(b []byte)(int,error){
	length:=len(b)
	return length,nil
}