package models

import "github.com/Shu682682/Booking.git/internal/forms"

//TemplateData holds data sent from hanlders to templates
type TemplateData struct{
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string] float32
	Data map[string]interface{}//not sure data type just use interface
	CSRFToken string
	Flash string
	Warning string
	Error string
	Form  *forms.Form


}
