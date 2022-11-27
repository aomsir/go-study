package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Name    string
	Address string
}

func main() {
	a := 1.5
	fmt.Println(reflect.TypeOf(a))  // float64
	fmt.Println(reflect.ValueOf(a)) // 1.5

	// 获取值
	people := People{"Aomsir", "Wuhan"}
	fmt.Println(reflect.TypeOf(people))                        // main.people
	fmt.Println(reflect.ValueOf(people))                       // {Aomsir Wuhan}
	fmt.Println(reflect.TypeOf(people).NumField())             // 2
	fmt.Println(reflect.TypeOf(people).FieldByIndex([]int{1})) //{address main string  16 [1] false}

	// 设置值
	// 反射时获取peo的地址
	// Elem()获取指针指向地址的封装
	// 地址的值必须调用Elem()才可以继续操作
	content := "Name"
	peo := new(People)
	peo.Name = "Aomsir"
	v := reflect.ValueOf(peo).Elem()

	// 需要修改属性的内容时,要求结构体中属性首字母大写才能
	fmt.Println(v.FieldByName(content).CanSet()) //true
	v.FieldByName(content).SetString("Jewix")
	fmt.Println(*peo) // {address main string  16 [1] false}

}
