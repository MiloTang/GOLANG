package action

import (
	"fmt"
	"net/http"
)

func Admin(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "这是管理页面!")
}
