package action

import (
	"fmt"
	"html/template"
	"net/http"
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
	Rst := mysql.Select("select * from guests")
	fmt.Println(Rst)
	p.Lists = Rst
	p.Title = "benoob ceshi"
	p.Details = "打算科技防盗锁"
	t.Execute(rw, p)
}
