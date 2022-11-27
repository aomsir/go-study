package main

import "fmt"

func main() {
	// 实例化的两种方式
	m := make(map[string]string)
	n := map[string]string{
		"name": "Aomsir",
		"age":  "20",
	}

	n["age"] = "21"               // key存在则覆盖
	n["email"] = "info@say521.cn" // key不存在则新增

	delete(n, "email") // key存在则删除
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(n["name"]) // key存在则返回对应的value,不存在就返回value类型的默认值

	// 查看key的值以及是否存在
	value, ok := n["name"]
	fmt.Println(value, ok)

	// 遍历所有的值
	for key, value := range n {
		fmt.Println(key, value)
	}
}
