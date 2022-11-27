package main

import "fmt"

func main() {
	// 单个defer
	// 执行结果： 打开连接->执行功能->关闭连接->关闭连接
	// fmt.Println("打开连接")
	// defer fmt.Println("关闭连接")
	// defer func() {
	// 使用匿名函数,和上面的等价
	// 	fmt.Println("关闭连接")
	// }()
	// fmt.Println("执行功能")

	// 多个defer,栈结果
	// 执行结果： 打开A连接->打开B连接->关闭b连接->关闭A连接
	fmt.Println("打开A连接")
	defer fmt.Println("关闭A连接")
	fmt.Println("打开B连接")
	defer fmt.Println("关闭B连接")
}
