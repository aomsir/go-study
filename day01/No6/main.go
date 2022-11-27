package main

import "fmt"

func main() {
	// 常量生成器
	const (
		a = iota
		b
		c
	)

	// 0 1 2
	fmt.Println(a, b, c)
}
