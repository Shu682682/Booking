package forms

import (
	"net/http"
	"net/url"
)

//From creates a custom from strut, embeds a url.Values object
type Form struct{
	url.Values
	Errors errors
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
	x:=r.Form.Get(field)
	if x==""{
		return false
	}
	return true
}