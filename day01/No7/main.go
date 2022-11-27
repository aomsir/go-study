package main

import "fmt"

func main() {
	var a = 1
	var b *int = &a

	*b = 2

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
	fmt.Println(*b) // 取变量内的具体内容
}
