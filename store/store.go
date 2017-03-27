package main

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func init() {

}

type Student struct {
	name string
	age  int
	id   int
}

var Students map[int]*Student

func memoryIn(st Student) {
	Students[st.id] = &st
}
func memoryOut() {
	for _, v := range Students {
		fmt.Println(v)
	}
}
func fileIn() {
	data := []byte("Hello World!")
	err := ioutil.WriteFile("stdata", data, 0664)
	if err != nil {
		panic(err)
	}
}
func fileOut() {
	rdata, err := ioutil.ReadFile("stdata")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(rdata))
}
func csvIn() {
	csvFile, err := os.Create("st.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	for _, st := range Students {
		line := []string{st.name, strconv.Itoa(st.age), strconv.Itoa(st.id)}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
}
func csvOut() {
	csvFile, err := os.Open("st.csv")
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
	fmt.Println(records)
}
func gobIn() {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode("汉字")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("gob", buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
func main() {
	Students = make(map[int]*Student)
	for i := 0; i < 99; i++ {
		memoryIn(Student{name: "milo", age: 20, id: i})
	}
	memoryOut()
	fileIn()
	fileOut()
	csvIn()
	csvOut()
	gobIn()
}
