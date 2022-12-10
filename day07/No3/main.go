package main

import "fmt"

type Student struct {
	name string `json:"name"`
	age  int    `json:"age"`
}

func main() {
	student := Student{"Aomsir", 21}
	fmt.Println(student)

}
