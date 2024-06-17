package main

import (
	"fmt"
	"unsafe"
)

func main1() {
	var num int
	var num8 int8
	var num16 int16
	var num32 int32
	var num64 int64

	fmt.Println("bum bytes: ", unsafe.Sizeof(num))
	fmt.Println("bum bytes: ", unsafe.Sizeof(num8))
	fmt.Println("bum bytes: ", unsafe.Sizeof(num16))
	fmt.Println("bum bytes: ", unsafe.Sizeof(num32))
	fmt.Println("bum bytes: ", unsafe.Sizeof(num64))
}