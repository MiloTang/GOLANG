package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

func init() {

}
func main() {
	var in bytes.Buffer
	b := []byte(`www.mogupao.com`)
	w := zlib.NewWriter(&in)
	zlib.NewWriterLevel(w, 10)
	w.Write(b)
	w.Close()
	fmt.Println(len(in.String()))
	var out bytes.Buffer
	r, _ := zlib.NewReader(&in)
	io.Copy(&out, r)
	fmt.Println(len(out.String()))
}
