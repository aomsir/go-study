package main

import "fmt"

func main() {
	i := 6
	if i == 6 {
		goto Loop
	}
	fmt.Println("执行完if")

Loop:
	fmt.Println("loop")
	fmt.Println("程序结束")
}
