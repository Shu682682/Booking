package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

//From creates a custom from strut, embeds a url.Values object
type Form struct{
	url.Values
	Errors errors
}

func (f *Form) Valid()bool{
	return len(f.Errors)==0
}
//New initializes a form struct
func New(data url.Values) *Form{
	return &Form{
		Values: data,
		Errors: make(errors),
	}
}

//Has checks if form field is in post and not empty
func (f *Form ) Has(field string, r *http.Request) bool{
	return f.Get(field) != ""
}

//Required checks for required fields

func (f *Form) Required(fields ...string){
	for _, field:=range fields{
		value:=f.Get(field)
		if strings.TrimSpace(value)==""{
			f.Errors.Add(field, "This field cannot be blank")
		}
		
	}
}


func (f*Form) MinLength(field string , length int, r *http.Request)bool{
	x:=f.Get(field)
	if len(x)<length{
		f.Errors.Add(field, fmt.Sprintf("This field must be at least #{length} characters long"))
		return false
	}
	return true
}

//check for valid email address

// func (f *Form) IsEmail(field string){
// 	if !govalidator.IsEmail(f.Get(field)){
// 		f.Errors.Add(field, "Invalid email address")
// 	}
// }