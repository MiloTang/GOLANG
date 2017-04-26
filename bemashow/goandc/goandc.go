package main

/*
#include <stdio.h>

void sayHi() {
    printf("Hi");
}
*/
import "C"

func init() {

}

func main() {
	C.sayHi()
}
