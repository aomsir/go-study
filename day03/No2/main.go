package main

import (
	"errors"
	"fmt"
)

func main() {
	result, error := demo(6, 2)
	fmt.Println(result, error) // 3 <nil>

	result, error = demo(6, 0)
	fmt.Println(result, error) // 0 除数不能为0
}

func demo(i, k int) (r int, e error) {
	if k == 0 {
		e = errors.New("除数不能为0")
		return
	} else {
		r = i / k
		return
	}
}
