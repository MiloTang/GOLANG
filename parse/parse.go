package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	err    error
	record []string
)

func main() {
	defer func() {
		if err1 := recover(); err1 != nil {
			fmt.Println(err1)
		}
	}()
	Csv()
	Json()
	Xml()
}
func Xml() {
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName   xml.Name `xml:"person"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address   Address
		Comment   string `xml:"comment"`
	}
	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write(output)

	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}
	value := Result{Name: "none", Phone: "none"}

	data := `
    <Person>
        <FullName>Grace R. Emlin</FullName>
        <Company>Example Inc.</Company>
        <Email where="home">
            <Addr>gre@example.com</Addr>
        </Email>
        <Email where='work'>
            <Addr>gre@work.com</Addr>
        </Email>
        <Group>
            <Value>Friends</Value>
            <Value>Squash</Value>
        </Group>
        <City>Hanga Roa</City>
        <State>Easter Island</State>
    </Person>
`
	err = xml.Unmarshal([]byte(data), &value)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", value.XMLName)
	fmt.Printf("Name: %q\n", value.Name)
	fmt.Printf("Phone: %q\n", value.Phone)
	fmt.Printf("Email: %v\n", value.Email)
	fmt.Printf("Groups: %v\n", value.Groups)
	fmt.Printf("Address: %v\n", value.Address)

}
func Csv() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))
	w := csv.NewWriter(os.Stdout)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		err = w.Write(record)
		CheckErr(err)
		fmt.Println(record)
	}
	w.Flush()

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
	wr := csv.NewWriter(os.Stdout)
	wr.WriteAll(records)
	if err = wr.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}
func Json() {
	type Student struct {
		Name string
		Age  int
	}
	st := []Student{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	b, err1 := json.Marshal(st)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(st)
	fmt.Println(string(b))
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	out.WriteTo(os.Stdout)
	var ust []Student
	err = json.Unmarshal(b, &ust)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("\n%+v\n", ust)
	for _, v := range ust {
		fmt.Println("名字是:", v.Name)
		fmt.Println("年龄是:", v.Age)
	}

	const jsonStream = `
    {"Name": "Ed", "Text": "Knock knock."}
    {"Name": "Sam", "Text": "Who's there?"}
    {"Name": "Ed", "Text": "Go fmt."}
    {"Name": "Sam", "Text": "Go fmt who?"}
    {"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}
