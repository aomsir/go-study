package main

import "fmt"

func main() {

abc:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 1 {
				continue abc
			}
			fmt.Println(i, j)
		}
	}
}
