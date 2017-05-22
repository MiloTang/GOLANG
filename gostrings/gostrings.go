package main

import (
	fm "fmt"
	"strings"
)

func init() {

}
func main() {
	var str string = "this is a sample strings!"
	String(str)
}
func String(str string) {
	fm.Println(strings.HasPrefix(str, "t"))
	fm.Println(strings.HasSuffix(str, "!"))
	fm.Println(strings.Count(str, "s"))
	fm.Println(strings.Index(str, "i"))
	fm.Println(strings.LastIndex(str, "s"))
	fm.Println(strings.NewReader(str))
}
