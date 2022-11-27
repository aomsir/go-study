package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "hello\nworld"
	fmt.Println(s)
	fmt.Println(s + "3")

	// _未变量,接收第二个返回值的,但没有用途
	i, _ := strconv.ParseInt("11", 2, 64)
	fmt.Println(i)

	s1 := strings.LastIndex(s, "d")
	fmt.Println(s1)

}
