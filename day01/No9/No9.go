package main

import "fmt"

func main() {
	score := 60
	if score >= 60 {
		fmt.Println("Hello")
	}

	if name := "Aomsir"; name == "Aomsir" {
		fmt.Println("Hello," + name)
	} else if name == "Jewix" {
		fmt.Println("Hello," + name)
	} else {
		fmt.Println("Hello," + "root")
	}
}
