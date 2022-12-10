package main

import "fmt"

func main() {
	arr := make([]int, 0, 0)
	arr = append(arr, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(arr)
	fmt.Println("---------------------------------")

	names := []string{"Aomsir", "Jewix"}
	names = append(names, "Hello")
	fmt.Println(names)
	fmt.Println("---------------------------------")

	array := [5]int{1, 2, 3, 4, 5}
	slice := array[0:4] // 切片去引用数组,0开始,前包后不包
	fmt.Println(slice)
}
