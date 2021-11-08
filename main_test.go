package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var app App

func TestMetricsHandler(t *testing.T) {
	app.Initialize()

	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	app.Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := "myapp_processed_ops_total"

	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestFooHandler(t *testing.T) {
	app.Initialize()

	req, err := http.NewRequest("GET", "/foo", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	app.Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"name":"Bar"},{"name":"Baz"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
func TestBarHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/bar", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	app.Initialize()
	app.Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"name":"Bar 1","addresses":["street 1","street 2"]},{"name":"Bar 2","addresses":["street 3","street 4"]}]`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
