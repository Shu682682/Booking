package main

import (
	"fmt"
	"testing"

	"github.com/Shu682682/Booking.git/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T){
	var app config.AppConfig
	mux:=routes(&app)
	switch v:=mux.(type){
	case *chi.Mux:
	default:
		t.Error(fmt.Sprintf("Type is not *chu.Mux, type is %T",v))

	}
}