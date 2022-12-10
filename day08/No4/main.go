package main

import (
	"fmt"
	"os"
	"strconv"
)

// go run main.go 3 4
// 7
func main() {
	sum := sum(os.Args[1], os.Args[2])
	fmt.Println(sum)
}

func sum(number1 string, number2 string) int{
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)

	return int1 + int2
}