package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcome, visit sub pages now."))
	})
	mux.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("data"))
	})
	mux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("json"))
	})
	mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
		// Assumes you have a template in ./templates called "example.tmpl"
		// $ mkdir -p templates && echo "<h1>Hello {{.}}.</h1>" > templates/example.tmpl
		w.Write([]byte("html"))
	})
	http.ListenAndServe(":8080", mux)
}
