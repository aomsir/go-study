package main

import "fmt"

func main() {
	f := closure() // f是返回值函数
	fmt.Println(f())
}

func closure() func() int {
	i := 1
	return func() int {
		i += 1
		return i
	}
}
