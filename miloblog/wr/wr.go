package wr

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title    string
	Lists    []Context
	Next     string
	Previous string
	Token    string
	Info     string
	Details  string
	Username string
	Password string
}

var (
	p = &Page{}
)

type Context struct {
	Introduction string
	Link         string
}

func init() {

}
func Life(w http.ResponseWriter, r *http.Request) {
	titles := []Context{}
	p.Title = "Milo Blog"
	p.Lists = titles
	t, _ := template.ParseFiles("life.html")
	t.Execute(w, p)
}

func Manual(w http.ResponseWriter, r *http.Request) {
	titles := []Context{}
	p.Title = "Milo Blog"
	p.Lists = titles
	t, _ := template.ParseFiles("manual.html")
	t.Execute(w, p)
}

func Api(w http.ResponseWriter, r *http.Request) {

}

func Index(w http.ResponseWriter, r *http.Request) {
	titles := []Context{}
	p.Title = "Milo Blog"
	p.Lists = titles
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func Admin(w http.ResponseWriter, r *http.Request) {
	titles := []Context{}
	p.Title = "Milo Blog"
	p.Lists = titles
	t, _ := template.ParseFiles("login.html")
	t.Execute(w, p)
}
