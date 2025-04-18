package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)


func TestForm_Valid(t *testing.T){
	r:=httptest.NewRequest("POST","/whatever",nil)
	form :=New(r.PostForm)
	isValid:=form.Valid()
	if !isValid{
		t.Error("got invalid when should have been valid")
	}
	
}

func TestFrom_Required(t *testing.T){
	r:=httptest.NewRequest("POST","/whatever",nil)

	form:=New(r.PostForm)
	form.Required("a","b","c")
	if form.Valid(){
		t.Error("from shows valid when required fields missing")
	}
	postedData:=url.Values{}
	postedData.Add("a","a")
	postedData.Add("b","a")
	postedData.Add("c","a")
	r,_=http.NewRequest("POST","/whatever",nil)
	r.PostForm=postedData
	form=New(r.PostForm)
	form.Required("a","b","c")
	if !form.Valid(){
		t.Error("shows does not have required fields when id does")
	}
}


func TestForm_Has(t *testing.T){
	r:=httptest.NewRequest("POST","/whatever",nil)
	form:=New(r.PostForm)
	has:=form.Has("whatever",r)
	if has{
		t.Error("form shows has field when it does not")
	}
	postedData:=url.Values{}
	postedData.Add("a","a")
	form=New(postedData)
	has=form.Has("a",r)
	if !has{
		t.Error("shows form does not have field when it should")
	}
}

func Test_MinLength(t *testing.T){
	r:=httptest.NewRequest("POST","/whatever",nil)
	form:=New(r.PostForm)
	form.MinLength("x",10,r)
	if form.Valid(){
		t.Error("form shows min length for non-existent field")
	}
	postedValues:=url.Values{}
	postedValues.Add("some_field","some value")
	form=New(postedValues)
	form.MinLength("some_field",100,r)
	if form.Valid(){
		t.Error("shows minlength  of 100 met when data is shorter")
	}
	postedValues=url.Values{}
	postedValues.Add("another_field","abc123")
	form=New(postedValues)
	form.MinLength("another_field",1,r)
	if !form.Valid(){
		t.Error("shows minlength of 1 is not met when it is")
	}
}