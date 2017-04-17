package main

/*
#include <stdio.h>

viod hello(){
	printf("hello go and C world ");
}
*/
import "C"

func init() {

}

func main() {
	C.hello()
}
