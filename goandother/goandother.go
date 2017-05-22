package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"
import (
	"fmt"
)

func main() {
	fmt.Println(C.int)
}
