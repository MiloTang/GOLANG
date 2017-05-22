package main

import (
	"fmt"
	"strconv"

	"github.com/axgle/mahonia"
)

var (
	a map[string]string
)

func init() {

}
func main() {

	a = make(map[string]string)
	a["b"] = "c" //这样才不会错
	hello := "你好世界"
	fmt.Println([]byte(hello))
	gbk := mahonia.NewDecoder("gbk").ConvertString(hello)
	fmt.Println(gbk)
	fmt.Println([]byte(gbk))
	utf8 := mahonia.NewEncoder("gbk").ConvertString(gbk)
	fmt.Println(utf8)
	fmt.Println(f(9)("h"))
	fmt.Println(f(9)("d"))
}
func f(i int) func(str string) string {
	return func(str string) string {
		i++
		return strconv.Itoa(i) + str
	}
}
