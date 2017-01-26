package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Page struct {
	Title string
	Lists []Context
}
type Context struct {
	Introduction string
	Link         string
}

func index(w http.ResponseWriter, r *http.Request) {
	p := &Page{}
	cs := []Context{{"Go表单", "/goform/"}, {"Go令牌", "/fotoken/"}}
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
func goform(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
func main() {
	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/index", index)
	http.HandleFunc("/goform", goform)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
