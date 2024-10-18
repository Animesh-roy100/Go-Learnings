package main

import (
	"fmt"
	"unsafe"
)

// BadStruct: Inefficient field ordering
type BadStruct struct {
	a bool  // 1 byte
	b int64 // 8 bytes
	c bool  // 1 byte
}

// GoodStruct: Efficient field ordering
type GoodStruct struct {
	b int64 // 8 bytes
	a bool  // 1 byte
	c bool  // 1 byte
}

func main() {
	fmt.Printf("Size of BadStruct: %d bytes\n", unsafe.Sizeof(BadStruct{}))
	fmt.Printf("Size of GoodStruct: %d bytes\n", unsafe.Sizeof(GoodStruct{}))
}
