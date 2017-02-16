package main

import (
	"crypto/md5"
	"fmt"
	"goweb/golib/gocs"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
	err   error               = nil
	file  *os.File
)

func index(w http.ResponseWriter, r *http.Request) {
	titles := []Context{}
	titles = bloglists(titles, "blog")
	p.Title = "Milo Blog"
	p.Lists = titles
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}
func onboot(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("bootstrap.html")
	t.Execute(w, nil)
}
func layoutit(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layoutit.html")
	t.Execute(w, nil)
}
func editor(w http.ResponseWriter, r *http.Request) {
	csid := startcs(w, r, "editor")
	_, b := cs.GetSession(csid, "username")
	if !b {
		http.Redirect(w, r, "/formlogin/", 302)
	}
	p.Title = "blog编辑"
	if r.Method == "GET" {
		t, _ := template.ParseFiles("editor.html")
		t.Execute(w, p)
	} else {
		r.ParseForm()
		title := r.FormValue("title")
		context := r.FormValue("content")
		if len(title) < 10 {
			io.WriteString(w, "标题太短了")
			return
		}
		if len(context) < 100 {
			io.WriteString(w, "少于100字的也算是文章？")
			return
		}
		err = save("blog/"+title, context)
		if err != nil {
			fmt.Println(err)
			io.WriteString(w, "出错了")
		} else {
			io.WriteString(w, "true")
		}
	}
}
func delsession(w http.ResponseWriter, r *http.Request) {
	cs.DestroySession(w, r)
	http.Redirect(w, r, "/index/", 302)
}
func ErrorPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("error.html")
	t.Execute(w, p)
}
func save(filename string, context string) error {
	_, err = os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(filename)
			if err != nil {
				log.Println(err)
				return err
			}
		} else {
			log.Println(err)
			return err
		}
	} else {
		file, err = os.OpenFile(filename, os.O_WRONLY, 0666)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	defer file.Close()
	_, err = file.WriteString(context)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func startcs(w http.ResponseWriter, r *http.Request, function string) string {
	var csid string = ""
	_, err = r.Cookie(gocs.CookieName)
	if err != nil {
		if err.Error() == "http: named cookie not present" {
			csid, err = cs.StartSession(w, r)
			if err != nil {
				panic(err)
			} else {
				http.Redirect(w, r, "/"+function+"/", 302)
			}
		} else {
			panic(err)
		}
	} else {
		csid, err = cs.StartSession(w, r)
		if err != nil {
			panic(err)
		}
	}
	return csid
}
func details(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	spurls := strings.Split(url, "/")
	t, _ := template.ParseFiles("details.html")
	p.Title = spurls[2]
	dets, _ := read("blog/" + spurls[2])
	p.Details = dets
	for i, _ := range p.Lists {
		if spurls[2] == p.Lists[i].Introduction {
			if i > 0 {
				p.Previous = "/" + p.Lists[i-1].Link
			} else {
				p.Previous = "#"
			}
			if i < (len(p.Lists) - 1) {
				p.Next = "/" + p.Lists[i+1].Link
			} else {
				p.Next = "#"
			}
		}
	}
	t.Execute(w, p)
}
func read(filename string) (string, error) {
	var context []byte
	_, err = os.Stat(filename)
	if err != nil {
		log.Println(err)
		return " ", err
	} else {
		file, err = os.Open(filename)
		if err != nil {
			log.Println(err)
			return " ", err
		}
		defer file.Close()
		context, err = ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
			return " ", err
		}
	}
	return string(context), nil
}
func bloglists(titles []Context, path string) []Context {
	lists, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		return titles
	}
	PthSep := string(os.PathSeparator)
	for _, ls := range lists {
		if ls.IsDir() {
			bloglists(titles, path+PthSep+ls.Name())
		} else {
			titles = append(titles, Context{ls.Name(), "details" + "/" + ls.Name()})
		}
	}
	return titles
}
func formlogin(w http.ResponseWriter, r *http.Request) {
	csid := startcs(w, r, "formlogin")
	_, b := cs.GetSession(csid, "username")
	if b {
		http.Redirect(w, r, "/editor/", 302)
	}
	p.Title = "Milo Blog Login"
	p.Username = ""
	if r.Method == "GET" {
		cs.SetSession(csid, "token", token())
		p.Info = ""
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, p)
	} else {
		r.ParseForm()
		username := template.HTMLEscapeString(r.FormValue("username"))
		password := template.HTMLEscapeString(r.FormValue("password"))
		tk := template.HTMLEscapeString(r.FormValue("token"))
		cstk, _ := cs.GetSession(csid, "token")
		if len(username) == 0 {
			p.Info = "用户名不能为空"
			p.Username = username
			p.Password = password
			cs.SetSession(csid, "token", token())
			t, _ := template.ParseFiles("login.html")
			t.Execute(w, p)
			return
		}
		if len(password) < 6 {
			p.Info = "密码不能小于6位"
			p.Username = username
			p.Password = password
			cs.SetSession(csid, "token", token())
			t, _ := template.ParseFiles("login.html")
			t.Execute(w, p)
			return
		}
		_, err = os.Stat("user/" + username)
		if err != nil {
			if os.IsNotExist(err) {
				p.Info = "用户名不存在"
				p.Username = username
				p.Password = password
				cs.SetSession(csid, "token", token())
				t, _ := template.ParseFiles("login.html")
				t.Execute(w, p)
				return
			}
			panic(err)
		} else {
			ps, e := ioutil.ReadFile("user/" + username)
			if e != nil {
				panic(e)
			}
			if tomd5(password) == string(ps) {
				if tk == cstk {
					cs.SetSession(csid, "token", "")
				} else {
					io.WriteString(w, "不要重复提交")
					return
				}
				cs.SetSession(csid, "username", username)
				http.Redirect(w, r, "/editor/", 302)
			} else {
				p.Info = "密码不对"
				p.Username = username
				p.Password = password
				cs.SetSession(csid, "token", token())
				t, _ := template.ParseFiles("login.html")
				t.Execute(w, p)
				return
			}
		}
	}
}

func token() string {
	crutime := time.Now().Unix()
	t := md5.New()
	io.WriteString(t, strconv.FormatInt(crutime, 10))
	p.Token = fmt.Sprintf("%x", t.Sum(nil))
	return p.Token
}
func tomd5(str string) string {
	md := md5.New()
	md.Write([]byte(str))
	return fmt.Sprintf("%x", md.Sum(nil))
}
func main() {
	gocs.MaxLT = 3600
	gocs.CookieName = "Miloblog"
	cs = gocs.NewCookieSession()
	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))
	http.Handle("/bootstrap/", http.FileServer(http.Dir("static")))
	http.Handle("/layoutitlib/", http.FileServer(http.Dir("static")))
	http.Handle("/wysiwyg/", http.FileServer(http.Dir("static")))
	http.Handle("/fonts/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/", index)
	http.HandleFunc("/details/", details)
	http.HandleFunc("/formlogin/", formlogin)
	http.HandleFunc("/delsession/", delsession)
	http.HandleFunc("/editor/", editor)
	http.HandleFunc("/onboot/", onboot)
	http.HandleFunc("/layoutit/", layoutit)
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
		log.Println(e.Error())
	}
}
