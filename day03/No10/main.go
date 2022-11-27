package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("Hello,Aomsir")
	b := make([]byte, r.Size())
	length, error := r.Read(b)

	if error != nil {
		fmt.Println("读取数据流失败", error)
		return
	}
	fmt.Println("读取数据的长度为", length)
	fmt.Println("数据内容为", string(b))
}
