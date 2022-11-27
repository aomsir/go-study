package main

import "fmt"

func main() {
	names := []string{"1", "2", "3", "4"} //since嗷,没有定义大小
	names[4] = "5"
	fmt.Println(names[4])
}
