package main

import (
	"fmt"
	"study/day02/No8/src/people"
)

func main() {
	people := new(people.People)
	people.SetName("Aomsir")
	people.SetAge(20)
	fmt.Println(*people)
}
