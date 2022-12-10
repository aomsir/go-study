package main

import (
	"fmt"
	"study/day08/No6/calculator"
)

func main() {
    total := calculator.Sum(3, 5)
    fmt.Println(total)
    fmt.Println("Version: ", calculator.Version)
}