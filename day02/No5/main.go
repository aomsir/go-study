package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

func main() {
	var animal Animal

	// 第一种赋值
	animal.Name = "Aomsir"
	animal.Age = 20

	// 第二种赋值
	animal = Animal{"Aomsir", 20}

	// 第三种赋值
	animal1 := Animal{"Aomsir", 20}
	fmt.Println(animal)
	fmt.Println(animal1)

	if animal1 == animal {
		fmt.Println("两者内容相同")
	}
}
