package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{2, 1, 4, 8, 9}
	// sort.Sort(sort.IntSlice(num)) 与下面等价
	sort.Ints(num) // 升序排序
	fmt.Println(num)

	sort.Sort(sort.Reverse(sort.IntSlice(num))) // 降序排列
	fmt.Println(num)
}
