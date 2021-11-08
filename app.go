package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type App struct {
	Router *mux.Router
}

type Foo struct {
	Name string `json:"name"`
}

type Bar struct {
	Name      string   `json:"name"`
	Addresses []string `json:"addresses"`
}

var (
	fooProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "foo_processed_ops_total",
		Help: "The total number of foo processed events",
	})
)

var (
	barProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bar_processed_ops_total",
		Help: "The total number of bar processed events",
	})
)

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/foo", a.getFoos).Methods("GET")
	a.Router.HandleFunc("/bar", a.getBars).Methods("GET")
	a.Router.Handle("/metrics", promhttp.Handler())
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) getFoos(w http.ResponseWriter, r *http.Request) {
	defer fooProcessed.Inc()

	foos := []Foo{
		{Name: "Bar"},
		{Name: "Baz"},
	}

	js, err := json.Marshal(foos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(js)
}

func (a *App) getBars(w http.ResponseWriter, r *http.Request) {
	defer barProcessed.Inc()

	bars := []Bar{
		{Name: "Bar 1", Addresses: []string{"street 1", "street 2"}},
		{Name: "Bar 2", Addresses: []string{"street 3", "street 4"}},
	}

	js, err := json.Marshal(bars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(js)
}
