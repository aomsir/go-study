package main

import "fmt"

func main() {
	a, b := 1, 2
	c := add(a, b)
	d := ded(a, b)
	fmt.Println(c)
	fmt.Println(d)
}

func add(a int, b int) int {
	return a + b
}

// 参数与函数的第二种写法
func ded(a, b int) (sum int) {
	sum = a + b
	return // 返回的是sum变量
}
