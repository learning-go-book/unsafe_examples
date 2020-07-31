package main

import (
	"fmt"
	"unsafe"
)

type Data struct {
	Value  uint32   // 4 bytes
	Label  [10]byte // 10 bytes
	Active bool     // 1 byte
	// padded with 1 byte to make it align
}

func DataFromBytesUnsafe(b [16]byte) Data {
	return *(*Data)(unsafe.Pointer(&b))
}

func DataFromBytes(b [16]byte) Data {
	d := Data{}
	d.Value = uint32(b[3])<<24 + uint32(b[2])<<16 + uint32(b[1])<<8 + uint32(b[0])
	copy(d.Label[:], b[4:14])
	d.Active = b[14] != 0
	return d
}

func BytesFromDataUnsafe(d Data) [16]byte {
	return *(*[16]byte)(unsafe.Pointer(&d))
}

func BytesFromData(d Data) [16]byte {
	out := [16]byte{}
	out[3] = byte(d.Value >> 24)
	out[2] = byte(d.Value >> 16)
	out[1] = byte(d.Value >> 8)
	out[0] = byte(d.Value & 255)
	copy(out[4:14], d.Label[:])
	if d.Active {
		out[14] = 1
	}
	return out
}

func main() {
	d := Data{
		Value:  8675309,
		Active: true,
	}
	copy(d.Label[:], "Phone")
	fmt.Println(d, unsafe.Alignof(d), unsafe.Alignof(d.Value), unsafe.Alignof(d.Label), unsafe.Alignof(d.Active))

	b := [16]byte{237, 95, 132, 0, 80, 104, 111, 110, 101, 0, 0, 0, 0, 0, 1, 0}
	fmt.Println(b)

	b1 := BytesFromData(d)
	b2 := BytesFromDataUnsafe(d)
	if b1 != b2 {
		panic(fmt.Sprintf("%v %v", b1, b2))
	}
	fmt.Printf("%+v\n", b1)
	d1 := DataFromBytes(b1)
	d2 := DataFromBytesUnsafe(b1)
	if d1 != d2 {
		panic(fmt.Sprintf("%v %v", d1, d2))
	}
	fmt.Println(d1)
}
