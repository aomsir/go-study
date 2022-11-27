package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3} // 会在内存中开辟一块连续的空间
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	arr2 := [...]int{1, 2, 3} // 不需要指定长度,有几个元素长度就是几

	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)
	arr2[0] = 9
	fmt.Println(arr2[0])
}
