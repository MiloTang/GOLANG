package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	SampleReadFormString()
	SampleReadStdin()
	SampleReadFile()
	SampleFileLine()
	SampleReadImage()
}

func ReadForm(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err

}
func SampleReadFormString() {
	data, _ := ReadForm(strings.NewReader("一些测试的字符串"), 20)
	strRead := strings.NewReader("测试缓冲字 符串")
	bufstr := bufio.NewReader(strRead)
	bufdata, _ := bufstr.Peek(6)
	bufsub, _ := bufstr.ReadString(' ')
	fmt.Println("读取字符串测试:", data)
	fmt.Println("缓冲字符串测试:", string(bufdata), bufstr.Buffered())
	fmt.Println("截取读", bufsub)
}
func SampleReadStdin() {
	fmt.Println("请输入:")
	data, _ := ReadForm(os.Stdin, 20)
	fmt.Println("你输入的 字符串是:", data, "===", string(data))
}
func SampleReadFile() {
	file, _ := os.Open("basicio.go")
	defer file.Close()
	data, _ := ReadForm(file, 200)
	fmt.Println(string(data))
}
func SampleFileLine() {
	if len(os.Args) < 2 {
		return
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	line := 0
	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
			break
		}
		if !isPrefix {
			line++
		}
	}
	fmt.Println(line)

}
func SampleReadImage() {
	file, err := os.Open("IMG.JPG")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	var head byte
	for {
		err := binary.Read(file, binary.LittleEndian, &head)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Print(head)
		fmt.Print(" ")
	}

}
