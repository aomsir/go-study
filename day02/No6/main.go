package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func main() {
	people := new(People) // 创建结构体类型指针
	fmt.Println(people)

	people1 := &People{"Aomsir", 20} // 创建结构体指针并赋值
	fmt.Println(people1)
}
