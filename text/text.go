package main

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

const N int16 = int16(unsafe.Sizeof(0))

var (
	sstr = "show me show me again,it's me"
	cstr = "me"
)

type person struct {
	Name string `xml:"name,attr"`
	Age  int    `xml:"age,attr"`
}

func main() {
	var x int16 = 0x1122
	p := unsafe.Pointer(&x)
	fmt.Println(p, N)
	p2 := (*[N]byte)(p)
	fmt.Println(p2[0], p2[1])
	if p2[0] == 17 {
		fmt.Println("本机器：大端")
	} else {
		fmt.Println("本机器：小端")
	}
	var xx int16
	var xa, xb byte
	xx = 0x1122
	xa = (*[N]byte)(unsafe.Pointer(&xx))[0]
	xb = (*[N]byte)(unsafe.Pointer(&xx))[1]
	fmt.Println(xa)
	fmt.Println(xb)
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
