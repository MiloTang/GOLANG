package main

import (
	"bemashow/webnet/action"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strings"
)

var (
	routes = map[string]func(rw http.ResponseWriter, req *http.Request){
		"/":      action.Index,
		"/admin": action.Admin,
	}
)

func init() {

}
func main() {
	http.HandleFunc("/", Entry)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func Entry(rw http.ResponseWriter, req *http.Request) {
	urls := strings.Split(req.URL.Path, "/")
	if handler, ok := routes["/"+urls[1]]; ok {
		handler(rw, req)
		return
	} else {
		fmt.Fprintln(rw, "请求页面不存在!")
	}
}
