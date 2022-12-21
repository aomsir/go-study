package main

import "fmt"

//func main() {
//	for i := 1; i <= 100; i++ {
//		if i/3 == 0 {
//			fmt.Println("Fizz")
//		} else if i/5 == 0 {
//			fmt.Println("Buzz")
//		} else if i/3 == 0 && i/5 == 0 {
//			fmt.Println("FizzBuzz")
//		} else {
//			fmt.Println(i)
//		}
//	}
//}

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		default:
			fmt.Println(i)
		}
	}
}
