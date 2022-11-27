package main

import "fmt"

func main() {
	// 匿名函数 - 无参数无返回值
	func() {
		fmt.Println("匿名函数 - 无参数无返回值")
	}() // 括号为调用

	// 匿名函数 - 有参数无返回值
	func(name string) {
		fmt.Println("name:" + name)
	}("Aomsir")

	// 匿名函数 - 无参数有返回值
	name := func() string {
		return "Aomsir"
	}()
	fmt.Println(name)
}
