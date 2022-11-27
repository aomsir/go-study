package main

import "fmt"

func main() {
	name, age, _ := show()
	fmt.Println(name, age)
}

func show() (string, int, string) {
	fmt.Println("执行了show()")
	return "Aomsir", 20, "info@say521.cn"
}
