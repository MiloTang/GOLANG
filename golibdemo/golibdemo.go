package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/gob"
	"fmt"
	"gotesting"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var (
	week                time.Duration
	firstname, lastname string
	content             string
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func init() {

}
func main() {
	uctest()
}
func uctest() {
	str1 := "USING package uc aaa"
	fmt.Println(gotesting.UpperCase(str1))
}

func TimeShow() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	fmt.Println(t.UTC())
	week = 60 * 60 * 24 * 7 * 1e9
	fmt.Println(t.Add(week))
	fmt.Println(t.Date())
	fmt.Println(t.Clock())
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Zone())
	fmt.Println(strings.Split(t.String(), " ")[1])
}
func InputShow() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstname, &lastname)
	fmt.Printf("Hi %s %s!\n", firstname, lastname)
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	instr, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("The input was: %s\n", instr)
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err = enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}

	hasher := sha1.New()
	io.WriteString(hasher, "test")
	b := []byte{}
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))
	//
	hasher.Reset()
	data := []byte("test")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)

}
