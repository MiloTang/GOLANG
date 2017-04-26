package main

import (
	"comlib/godb/gomysql"
	"comlib/gofile"
	"comlib/net/getbody"
	"fmt"
)

func init() {
}
func main() {
	gomysql.Connect()
	defer gomysql.Close()
	fmt.Println(gomysql.Select("select * from guests"))
	for k, v := range gomysql.Select("select * from guests") {
		fmt.Println(k, "==", v["reg_date"])
	}
}
func show() {

	fmt.Println(gofile.ListDirs("c:/Milo/boxs/"))
	fmt.Println(getbody.GetBody("http://www.cnblogs.com/golove/p/3269099.html"))
}
