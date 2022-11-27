package main

import "fmt"

func main() {
	var name, age string
	score := 100
	fmt.Println("请输入用户名和年龄:")
	// fmt.Scanln(&name, &age)
	fmt.Scanf("%s%s", &name, &age)
	fmt.Println("用户名为", name)
	fmt.Println("年龄为", age)
	fmt.Println("分数为", score)
}
