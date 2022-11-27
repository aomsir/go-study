package main

import "fmt"

func main() {
	defer fmt.Println("defer的执行")
	fmt.Println("1111111111")
	panic("遇到panic")
	fmt.Println("2222222222")
}
