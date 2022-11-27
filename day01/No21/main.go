package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	copy(s1, s2[0:2])

	fmt.Println(s1) // 4 5 3
}
