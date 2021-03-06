package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	err  error
	path string
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
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
			fmt.Println("dir:", path+PthSep+list.Name())
			ListDirs(path + PthSep + list.Name())
		} else {
			fmt.Println("file:", path+PthSep+list.Name())
		}
	}
}
