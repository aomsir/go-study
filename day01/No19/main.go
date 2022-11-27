package main

import "fmt"

func main() {
	s := make([]int, 0) // 初始长度为0,容量也是
	s = append(s, 1)
	fmt.Println(s[0])

	s1 := []int{2, 3, 4}
	s = append(s, s1...) // ...标明它是一个切片
	fmt.Println(s)

}
