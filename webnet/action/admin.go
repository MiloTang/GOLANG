package action

import (
	"html/template"
	"net/http"
)

func Admin(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("tpl/admin.htm")
	t.Execute(rw, nil)
}
