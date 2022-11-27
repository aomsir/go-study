package main

import (
	"fmt"
	"time"
)

func main() {
	// 声明时间变量
	var t time.Time
	fmt.Println(t)

	// 获取当前时间
	t1 := time.Now()
	fmt.Println(t1)

	// 通过纳秒时间戳创建
	t2 := time.Unix(0, t1.UnixNano())
	fmt.Println(t2)

	// 指定时间
	y, m, d := t2.Date()
	fmt.Println(y, m, d)
}
