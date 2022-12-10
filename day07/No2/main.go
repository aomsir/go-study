package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}
	arr1 := make([]int, 5)

	*(&arr[0]) = 100
	*(&arr1[0]) = 1000
	fmt.Println(arr)  // [100 2 3 4 5]
	fmt.Println(arr1) // [1000 2 3 4 5]

	fmt.Println(cap(arr1))
}
