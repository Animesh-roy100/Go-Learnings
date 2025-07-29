package main

/*
#include <stdio.h>

void sayHello() {
    printf("Hello from C!\n");
}

int add(int a, int b){
	return a + b;
}
*/
import "C"

func main() {
	C.sayHello()

	result := C.add(10, 20)
	println("result", result)
}
