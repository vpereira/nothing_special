package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Foo struct {
	Name string
}

type Bar struct {
	Name      string
	Addresses []string
}

func fakeMetric() {
	for {
		opsProcessed.Inc()
		time.Sleep(2 * time.Second)
	}
}
func recordMetrics() {
	go fakeMetric()
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	foo := Foo{"Bar"}
	js, err := json.Marshal(foo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(js)
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	bar := Bar{"Foo", []string{"street a", "street b"}}
	js, err := json.Marshal(bar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(js)
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/bar", barHandler)

	http.ListenAndServe(":31337", nil)
}