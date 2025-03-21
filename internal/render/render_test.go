package render

import (
	"net/http"
	"testing"

	"github.com/Shu682682/Booking.git/internal/models"
	"github.com/alexedwards/scs/v2"
)




func TestAddDefaultData(t *testing.T) {
	// ✅ 確保 `session` 已經初始化
	if session == nil {
		session = scs.New()
		testApp.Session = session
	}

	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Fatal("Failed to create session request:", err)
	}

	// ✅ 放入 session
	session.Put(r.Context(), "flash", "123")

	// ✅ 立即讀取確認數據是否存入
	flash := session.GetString(r.Context(), "flash")
	if flash != "123" {
		t.Fatalf("Session value not stored properly, expected '123' but got: '%s'", flash)
	}

	// ✅ 測試 `AddDefaultData()`
	result := AddDefaultData(&td, r)

	// ✅ 確保 `AddDefaultData()` 能夠讀取 session
	if result.Flash != "123" {
		t.Errorf("flash not found in session, expected '123' but got: '%s'", result.Flash)
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	// ✅ 確保 session 存在
	ctx, err := session.Load(r.Context(), "")
	if err != nil {
		return nil, err
	}

	r = r.WithContext(ctx) // ✅ 重新加載帶有 session 的 context
	return r, nil
}
