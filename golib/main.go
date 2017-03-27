package main

import (
	"fmt"
	"golib/godb/gomysql"
	"golib/gofile"
)

func init() {
}
func main() {
	gomysql.Connect()
	defer gomysql.Close()
	fmt.Println(gomysql.Select("select * from guests"))
	fmt.Println(gofile.ListDirs("c:/Milo/boxs/"))
}
