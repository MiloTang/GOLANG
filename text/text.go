package main

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	sstr = "show me show me again,it's me"
	cstr = "me"
)

type person struct {
	Name string `xml:"name,attr"`
	Age  int    `xml:"age,attr"`
}

func main() {
	Contains()
	Count()
	Index()
	Split()
	HasPrefix()
	ChangeType()
	SerXml()
}
func Contains() {
	fmt.Println(strings.Contains(sstr, cstr))
	fmt.Println(strings.Contains(sstr, "milo"))
}
func Count() {
	fmt.Println(strings.Count(sstr, cstr))
}
func Index() {
	fmt.Println(strings.Index(sstr, cstr))
	fmt.Println(strings.LastIndex(sstr, cstr))
}
func Split() {
	fmt.Println(strings.Split(sstr, cstr))
	fmt.Println(strings.Join(strings.Split(sstr, cstr), cstr))
}
func HasPrefix() {
	fmt.Println(strings.HasPrefix(sstr, cstr), strings.HasSuffix(sstr, cstr))
}
func ChangeType() {
	var num int = 10086
	str := strconv.Itoa(num)
	fmt.Println(str, reflect.TypeOf(str))
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(i, reflect.TypeOf(i))
	}
	fmt.Println(strconv.ParseBool("T"))
	fmt.Println(strconv.ParseFloat("3.14", 32))
	fmt.Println(strconv.ParseFloat("3.14", 64))
	fmt.Println(strconv.FormatInt(1234, 8))
	fmt.Println(strconv.FormatInt(1234, 10))

}
func SerXml() {
	p := person{"小花", 20}
	if data, err := xml.Marshal(p); err != nil {
		fmt.Println(err)
		return
	} else {
		p1 := new(person)
		if err1 := xml.Unmarshal(data, p1); err1 != nil {
			fmt.Println(err1)
		} else {
			fmt.Println(p1)
		}
		fmt.Println(string(data))
	}
	if data, err := xml.MarshalIndent(p, "", " "); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(data))
	}

}
