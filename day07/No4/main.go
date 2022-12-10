package main

import "fmt"

func main() {
	Map := map[string]int{"Aomsir": 21, "Jewix": 21}

	for key, value := range Map {
		fmt.Println(key, value)
	}
}
