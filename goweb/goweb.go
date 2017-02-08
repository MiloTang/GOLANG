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
}
type Context struct {
	Introduction string
	Link         string
}

var (
	p                        = &Page{}
	temp                     = map[string]string{"token": ""}
	CS   *gocs.CookieSession = nil
)

type method func(http.ResponseWriter, *http.Request)

func index(w http.ResponseWriter, r *http.Request) {
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
	u, ok := temp["username"]
	fmt.Println(u)
	if ok {
		p.Username = u
		t, _ := template.ParseFiles("vip.html")
		t.Execute(w, p)
	} else {
		http.Redirect(w, r, "/goformlogin/", 302)
	}
}
func gosession(w http.ResponseWriter, r *http.Request) {
	CS.DestroySession(w, r)
	//http.Redirect(w, r, "/index/", 302)

}
func goformlogin(w http.ResponseWriter, r *http.Request) {
	p.Title = "Go Login示例"
	p.Username = ""
	if r.Method == "GET" {
		p.Info = ""
		token()
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, p)
	} else {
		r.ParseForm()
		if temp["token"] != "" && temp["token"] == r.Form.Get("token") {
			temp["token"] = ""
		} else {
			fmt.Fprintf(w, "不要重复提交")
			return
		}
		if len(r.Form["username"][0]) == 0 {
			p.Info = "用户名不能为空"
			token()
			p.Password = r.Form.Get("password")
			t, _ := template.ParseFiles("login.html")
			t.Execute(w, p)
			return
		}
		if len(r.Form["password"][0]) < 6 {
			p.Info = "密码不能小于6位"
			token()
			p.Username = r.Form.Get("username")
			t, _ := template.ParseFiles("login.html")
			t.Execute(w, p)
			return
		}
		temp["username"] = r.Form.Get("username")
		http.Redirect(w, r, "/gosuccess/", 302)
	}
}

func token() {
	crutime := time.Now().Unix()
	t := md5.New()
	io.WriteString(t, strconv.FormatInt(crutime, 10))
	p.Token = fmt.Sprintf("%x", t.Sum(nil))
	temp["token"] = p.Token
}
func main() {
	gocs.MaxLT = 3600
	gocs.CookieName = "mytest"
	CS = gocs.NewCookieSession()
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
