package main

import "fmt"

func main() {
	a := new(int) // 这个a是 *int类型,在堆区创建一块内存
	fmt.Println(a)

	*a = 123
	fmt.Println(*a)
}
