package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5) // r即代表双向链表,又代表第一个元素
	r.Value = 0
	r.Next().Value = 1
	r.Next().Next().Value = 2
	// r.Next().Next().Next().Value = 3
	// r.Next().Next().Next().Next().Value = 4

	r.Prev().Value = 4
	r.Prev().Prev().Value = 3

	r.Unlink(2) // 把当前元素后面的两个删除

	// 0 1 2 3 4
	r.Do(func(i interface{}) {
		fmt.Println(i)
	})
}
