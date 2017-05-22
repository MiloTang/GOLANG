package main

import (
	"fmt"
	"regexp"
)

func init() {

}
func main() {
	match, _ := regexp.MatchString("s([a-z]+)ch", "seach")
	fmt.Println(match)
	r, _ := regexp.Compile("s([a-z]+)ch")
	fmt.Println(r.MatchString("seach China"))
	fmt.Println(r.FindString("seach punch sinch"))
	fmt.Println(r.FindStringIndex("seach punch sinch"))
	fmt.Println(r.FindStringSubmatch("seach punch sinch"))
	fmt.Println(r.FindStringSubmatchIndex("seach punch sinch"))
	fmt.Println(r.FindAllString("seach punch sinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("seach punch sinch", -1))
	fmt.Println(r.FindAllString("seach punch sinch", 2))
	fmt.Println(r.Match([]byte("seach punch sinch")))

	r = regexp.MustCompile("s([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a seach", "<fruit>"))

}
