package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	i := 23123
	fmt.Printf("%c\n", i)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Println(unsafe.Sizeof(i))
}
