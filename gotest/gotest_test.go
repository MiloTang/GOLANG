package gotest

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	got := hello()
	expect := "hello func in package hello."
	fmt.Println(got)
	if got != expect {
		t.Errorf("got [%s] expected [%s]", got, expect)
	}
}

func BenchmarkHello(b *testing.B) {

	for i := 0; i < b.N; i++ {
		hello()
	}
	fmt.Println(b.N)
}

func ExampleHello() {
	hl := hello()
	fmt.Println(hl)
}
