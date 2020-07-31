package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "hello"
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(sHdr.Len) // prints 5

	for i :=0;i< sHdr.Len; i++ {
		bp := *(*byte)(unsafe.Pointer(sHdr.Data+uintptr(i)))
		fmt.Print(string(bp))
	}
	fmt.Println()

	sHdr.Len = sHdr.Len+10
	fmt.Println(s)

	b := []byte("goodbye")
	bHdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sHdr.Data = bHdr.Data
	sHdr.Len = bHdr.Len
	fmt.Println(s)

	b[0] = 'x'
	fmt.Println(s)
}
