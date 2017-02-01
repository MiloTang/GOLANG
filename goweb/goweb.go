package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Page struct {
	Title string
	Lists []Context
	Token string
}
type Context struct {
	Introduction string
	Link         string
}

var (
	p = &Page{}
)

func index(w http.ResponseWriter, r *http.Request) {
	cs := []Context{{"Go表单登录示例", "/goformlogin/"}, {"Go令牌", "/fotoken/"}}
	cs = append(cs, Context{"GoCross", "/gocross/"})
	p.Title = "Go Web Samples"
	p.Lists = cs
	fmt.Println("path", r.URL.Path)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	t, _ := template.ParseFiles("index.html")
	fmt.Println(p)
	t.Execute(w, p)
}
func goformlogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	p.Title = "Go Login示例"
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		p.Token = fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, p)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))
	http.Handle("/fonts/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/index/", index)
	http.HandleFunc("/goformlogin/", goformlogin)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
