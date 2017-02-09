package main

import (
	"crypto/md5"
	"fmt"
	"goweb/golib/gocs"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Page struct {
	Title    string
	Lists    []Context
	Token    string
	Info     string
	Username string
	Password string
	Error    error
}
type Context struct {
	Introduction string
	Link         string
}

var (
	p                         = &Page{}
	cs    *gocs.CookieSession = nil
	Debug bool                = true
)

type method func(http.ResponseWriter, *http.Request)

func index(w http.ResponseWriter, r *http.Request) {
	cs.StartSession(w, r)
	fmt.Println(r.URL.Path)
	contexts := []Context{{"Go表单登录示例", "/goformlogin/"}, {"Go令牌", "/gotoken/"}, {"Go删除Session", "/gosession/"}}
	contexts = append(contexts, Context{"GoCross", "/gocross/"})
	p.Title = "Go Web Samples"
	p.Lists = contexts
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}
func gosuccess(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("vip.html")
	t.Execute(w, p)
}
func gosession(w http.ResponseWriter, r *http.Request) {
	cs.DestroySession(w, r)
	//http.Redirect(w, r, "/index/", 302)

}
func ErrorPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("vip.html")
	t.Execute(w, p)

}
func goformlogin(w http.ResponseWriter, r *http.Request) {
	p.Title = "Go Login示例"
	p.Username = ""
	id, err := cs.GetSessionID(r)
	if err != nil {
		ShowBug(w, err)
		return
	}
	cs.SetSession(id, "token", "temptoken")
	if r.Method == "GET" {
		p.Info = ""
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, p)
	} else {
		r.ParseForm()
		username := template.HTMLEscapeString(r.FormValue("username"))
		password := template.HTMLEscapeString(r.FormValue("password"))
		if len(username) == 0 {
			p.Info = "用户名不能为空"
			p.Password = r.Form.Get("password")
			t, _ := template.ParseFiles("login.html")
			t.Execute(w, p)
			return
		}
		if len(password) < 6 {
			p.Info = "密码不能小于6位"
			p.Username = r.Form.Get("username")
			t, _ := template.ParseFiles("login.html")
			t.Execute(w, p)
			return
		}
		http.Redirect(w, r, "/gosuccess/", 302)
	}
}

func token() string {
	crutime := time.Now().Unix()
	t := md5.New()
	io.WriteString(t, strconv.FormatInt(crutime, 10))
	p.Token = fmt.Sprintf("%x", t.Sum(nil))
	return p.Token
}
func main() {
	gocs.MaxLT = 3600
	gocs.CookieName = "mytest"
	cs = gocs.NewCookieSession()
	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))
	http.Handle("/fonts/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/", index)
	http.HandleFunc("/gosuccess/", gosuccess)
	http.HandleFunc("/goformlogin/", goformlogin)
	http.HandleFunc("/gosession/", gosession)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func ShowBug(w http.ResponseWriter, e error) {
	if Debug {
		fmt.Fprintf(w, e.Error())
	} else {
		fmt.Fprintf(w, "出了点小问题，请稍等")
	}
}
