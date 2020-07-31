package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i64 int64 = 1_000_000
	i64ptr := uintptr(unsafe.Pointer(&i64))
	bytes := *(*[unsafe.Sizeof(i64)]byte)(unsafe.Pointer(i64ptr))
	fmt.Println(bytes)
	i32 := *(*int32)(unsafe.Pointer(i64ptr))
	fmt.Println(i32)

	i64ptr += 4
	bytesShifted := *(*[unsafe.Sizeof(i64)]byte)(unsafe.Pointer(i64ptr))
	fmt.Println(bytesShifted)

	i32Shifted := *(*int32)(unsafe.Pointer(i64ptr))
	fmt.Println(i32Shifted)

	i64Shifted := *(*int64)(unsafe.Pointer(i64ptr))
	fmt.Println(i64Shifted)
}

