package main

import "fmt"

func main() {
	name := new(string)
	*name = "Aomsir"
	fmt.Println(*name)
	update(name)
	fmt.Println(*name)
}

func update(name *string) {
	*name = "Jewix"
}