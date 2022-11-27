package main

import "fmt"

func main() {
	i := demo()
	fmt.Println(i) // 值为0
	i = demo1()
	fmt.Println(i) // 值为2
}

func demo() int {
	i := 0
	defer func() {
		i += 2
	}()
	return i
}

func demo1() (z int) {
	i := 0
	defer func() {
		z = i + 2
	}()
	return i
}
