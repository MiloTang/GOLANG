package gofile

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	err   error
	lists []string
	datas map[interface{}]interface{}
)

func init() {
	lists = nil
	datas = make(map[interface{}]interface{})
}
func ListDirs(path string) []string {
	dirlist, err := ioutil.ReadDir(path)
	PthSep := string(os.PathSeparator)
	if err != nil {
		fmt.Println("read dir error->", err)
		return nil
	}
	for _, ls := range dirlist {
		if ls.IsDir() {
			lists = append(lists, ls.Name())
			ListDirs(path + PthSep + ls.Name())
		} else {
			lists = append(lists, ls.Name())
		}
	}
	return lists
}
func MemoryIn(k, v interface{}) {
	datas[k] = v
}
func MemoryOut() interface{} {
	return datas
}
func FileIn(data, filename string) {
	databyte := []byte(data)
	err := ioutil.WriteFile(filename, databyte, 0664)
	if err != nil {
		panic(err)
	}
}
func FileOut(filename string) string {
	rdata, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(rdata)
}
func CsvIn(filename string, records ...interface{}) {
	csvFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	for _, record := range records {
		err := writer.Write(record.([]string))
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
}
func CsvOut(filename string) [][]string {
	csvFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = 0
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	return records
}
