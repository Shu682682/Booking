package main

import "testing"

//cd cmd/web
//go test -v (for all)ß
//go test -cover
//go test -coverprofile=coverage.out && go tool cover -html=coverage.out
func TestRun(t *testing.T){
	err:=run()
	if err!=nil{
		t.Error("Failed run()")
	}
}