package main

import "fmt"

func main() {
	// 结果： 111111111 -》出现了panic,panic的信息为 出现了panic
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("出现了panic,panic的信息为", error)
		}
	}()
	fmt.Println("111111111")
	panic("出现了panic")
	fmt.Println("222222222")
}
