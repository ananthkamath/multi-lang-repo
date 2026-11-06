package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Greeting string `yaml:"greeting"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World! (Go)")
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	raw := []byte("greeting: hello from yaml\n")
	var cfg Config
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, cfg.Greeting)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler)
	r.HandleFunc("/config", configHandler)

	fmt.Println("go-app listening on port 8080")
	http.ListenAndServe(":8080", r)
}
