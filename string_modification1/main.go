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
		fmt.Println(string(bp))
	}
	sHdr.Len = sHdr.Len+10
	fmt.Println(s)

	bp := (*byte)(unsafe.Pointer(sHdr.Data+2))
	*bp = *bp + 1
	fmt.Println(s)
}
