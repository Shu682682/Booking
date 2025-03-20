package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct{
	key string
	value string
}

var theTests= [] struct{
	name string
	url string
	method string
	params []postData
	expectedStatusCode int
}{
{"home","/","GET",[]postData{}, http.StatusOK},
{"about","/about","GET",[]postData{}, http.StatusOK},
{"gq","/generals_quarter","GET",[]postData{}, http.StatusOK},
{"book","/book","GET",[]postData{}, http.StatusOK},
{"major","/majors_suite","GET",[]postData{}, http.StatusOK},
{"contact","/contact","GET",[]postData{}, http.StatusOK},
// POST test cases
{"post-book-json", "/book-json", "POST", []postData{
	{key: "start", value: "2021-01-01"},
	{key: "end", value: "2022-01-01"},
}, http.StatusOK},

{"post-book", "/book", "POST", []postData{
	{key: "start_date", value: "2025-01-01"},
	{key: "end_date", value: "2025-03-01"},
	{key: "full_name", value: "John Smith"},
	{key: "email", value: "123@gmail.com"},
	{key: "phone", value: "123456789"},
	{key: "people_amount", value:"3"},
	{key: "room_choice", value: "generals' quarter"},
}, http.StatusOK},
}



func TestHandlers(t *testing.T){
	routes:=getRoutes()
	ts :=httptest.NewTLSServer(routes)
	defer ts.Close()
	for _,e:=range theTests{
		if e.method =="Get"{
			resp, err:=ts.Client().Get(ts.URL+e.url)
			if err!=nil{
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode!=e.expectedStatusCode{
				t.Errorf("for %s, expected %d but hgot %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}

		}else if e.method=="POST"{
			values := url.Values{}
			for _, x:=range e.params{
				values.Add(x.key, x.value)
			}
			resp, err:=ts.Client().PostForm(ts.URL + e.url, values)
			if err!=nil{
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode!=e.expectedStatusCode{
				t.Errorf("for %s, expected %d but hgot %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}


		}else if e.method == "POST_JSON" {
			jsonData := map[string]string{}
			for _, x := range e.params {
				jsonData[x.key] = x.value
			}
		
			jsonBody, _ := json.Marshal(jsonData)
			req, _ := http.NewRequest("POST", ts.URL+e.url, bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
		
			resp, err := ts.Client().Do(req)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
		
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}