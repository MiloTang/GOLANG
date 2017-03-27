package main

import (
	"fmt"
	"golib/godb/gomysql"
)

func init() {
}
func main() {
	gomysql.Connect()
	defer gomysql.Close()
	fmt.Println(gomysql.Select("select * from guests"))
}
