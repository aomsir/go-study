package main

import "fmt"

type People struct {
	string // 匿名属性
	int
}

func main() {
	people := People{"Aomsir", 20}
	fmt.Println(people.string) // 调用属性类型
	fmt.Println(people.int)
}
