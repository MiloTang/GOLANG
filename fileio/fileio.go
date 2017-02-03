package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

}
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}
func test() {
	var data []byte
	var err error
	data = []byte(
		`
		绿杨芳草长亭路，
		年少抛人容易去。
		楼头残梦五更钟，
		花底离愁三月雨。
		无情不似多情苦，
		一寸还成千万缕。
		天涯地角有穷时，
		只有相思无尽处。`)
	err = ioutil.WriteFile("data1", data, 0644)
	CheckError(err)
	data, err = ioutil.ReadFile("data1")
	CheckError(err)
	fmt.Println(string(data))

	file, err := os.Create("data2")
	CheckError(err)
	defer file.Close()
	n2, err := file.WriteString("show me box\n")
	CheckError(err)
	fmt.Printf("wrote %d bytes\n", n2)
	n3, err := file.WriteString("writes agian\n")
	CheckError(err)
	fmt.Printf("wrote %d bytes\n", n3)
	d3 := []byte{115, 111, 109, 101, 10}
	d6 := []byte("show agian\n")
	n5, err := file.Write(d3)
	CheckError(err)
	fmt.Printf("wrote %d bytes\n", n5)
	n6, err := file.Write(d6)
	CheckError(err)
	fmt.Printf("wrote %d bytes\n", n6)
	file.Sync()
	w := bufio.NewWriter(file)
	n7, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n7)
	w.Flush()

	fileR, errr := os.Open("data2")
	defer fileR.Close()
	CheckError(errr)
	b1 := make([]byte, 1024)
	n1, errr := fileR.Read(b1)
	CheckError(errr)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))
}
