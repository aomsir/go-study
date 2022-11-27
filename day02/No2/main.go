package main

import "fmt"

func main() {
	demo("1", "2", "3", "4", "5")
}

func demo(hover ...string) {
	for i, n := range hover {
		fmt.Println(i, n)
	}
}
