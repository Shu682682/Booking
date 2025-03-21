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
	// 註冊 `gob` 以防止 session 無法序列化
	gob.Register(models.Reservation{})
	testApp.Session = session


	// 初始化 session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	// 設定 `testApp`
	testApp.InProduction = false
	testApp.Session = session

	// 建立 template cache
	tc, err := CreateTemplateCache() 
	if err != nil {
		log.Fatal("Cannot create template cache:", err)
	}

	testApp.TemplateCache = tc
	testApp.UseCache = false

	// 初始化 `render`
	NewTemplates(&testApp)

	// 運行測試
	exitCode := m.Run()

	// 退出
	os.Exit(exitCode)
}
