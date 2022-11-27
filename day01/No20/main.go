package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5}
	s := num[0:2] // 包前不包后,结果为1，2，s为切片
	fmt.Println(s)

	fmt.Printf("%p %p\n", num, &s[0]) // 地址一样

	s[0] = 6
	fmt.Println(num) // 数组的第一个元素也被改了
	fmt.Println(s)

	// 如果切片中添加数据,超过数组容量就会为切片重新分配空间
}
