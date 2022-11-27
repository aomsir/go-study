package main

import "fmt"

// 未运算
func main() {
	// 原码：0000 0011
	// 反码：0000 0011
	// 补码：0000 0011
	var a int8 = 3

	// 原码：0000 0010
	// 反码：1111 1101
	// 补码：1111 1110
	var b int8 = -2

	// 左移,正数乘2的几次方,操作的是原码
	fmt.Println(a << 2) // 0000 0011

	// 左移,负数乘2的几次方
	fmt.Println(b << 2) // 0000 1000

	fmt.Println(b >> 4) // 0000 1000
	fmt.Println(b >> 4) // 0000 1000

}
