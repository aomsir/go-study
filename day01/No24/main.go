package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 实例化list
	myList := list.New()
	fmt.Println(myList)

	fmt.Println("------------------------")

	// 添加元素
	myList.PushFront("a") // 头上加元素
	myList.PushBack("b")  // 尾巴上加元素
	myList.PushBack("c")

	myList.InsertBefore("d", myList.Back()) // 在最后一个元素的前添加
	fmt.Println(myList.Len())

	fmt.Println("------------------------")

	// 遍历
	for element := myList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println("------------------------")

	// 取第3个元素
	n := 3
	if n > 0 && n < myList.Len() {
		var curr *list.Element
		if n == 1 {
			curr = myList.Front()
		} else if n == myList.Len() {
			curr = myList.Back()
		} else {
			curr = myList.Front()
			for i := 1; i < n; i++ {
				curr = curr.Next()
			}
		}
		fmt.Println(curr.Value)
	} else {
		fmt.Println("n的数值不对")
	}

	fmt.Println("------------------------")

	// 移动元素的顺序
	myList.MoveBefore(myList.Front(), myList.Back()) // 将第一个移动到最后一个元素的前
	for element := myList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// 移除元素
	myList.Remove(myList.Front()) // 移除第一个元素
}
