package main

import "fmt"

func main() {
	for num := 0; num < 5; num++ {
		fmt.Println(num)
	}

	// 死循环
	for ; ; {
		fmt.Println("running...")
	}
}
