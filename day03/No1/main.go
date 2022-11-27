package main

import "fmt"

func main() {
	var i interface{} = 456
	result, response := i.(int)
	fmt.Println(result, response) // 456 true
	fmt.Printf("%T", result)      // int

	demo(456)   // 参数是int类型
	demo(456.9) // 参数是float类型
	demo(true)  // 参数类型不确定
}

func demo(i interface{}) {
	_, response := i.(int)
	if response == true {
		fmt.Println("参数是int类型")
		return
	}

	_, response = i.(float64)
	if response == true {
		fmt.Println("参数是float类型")
		return
	}

	fmt.Println("参数类型不确定")
	return
}
