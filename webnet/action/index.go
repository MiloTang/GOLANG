package action

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"webnet/comlib/db/mysql"
	"webnet/config"
)

type Page struct {
	Title   string
	Details string
	Lists   []map[string]interface{}
}

var (
	p = &Page{}
)

func init() {
}
func Index(rw http.ResponseWriter, req *http.Request) {
	config.CS.Start(rw, req)
	t, _ := template.ParseFiles("tpl/index.htm")
	mysql.Connect()
	defer mysql.Close()
	Rst := mysql.Select("select * from guests")
	fmt.Println(Rst)
	p.Lists = Rst
	p.Title = "benoob ceshi"
	p.Details = "打算科技防盗锁"
	t.Execute(rw, p)
}
func Register(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		t, _ := template.ParseFiles("tpl/register.htm")
		t.Execute(rw, nil)
	} else {
		req.ParseForm()
		username := template.HTMLEscapeString(req.FormValue("username"))
		email := template.HTMLEscapeString(req.FormValue("email"))
		agestr := template.HTMLEscapeString(req.FormValue("age"))
		age, _ := strconv.Atoi(agestr)
		mysql.Connect()
		defer mysql.Close()
		mysql.DML("insert into guests(username,email,age) values(?,?,?)", username, email, age)
	}
}
