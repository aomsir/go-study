package main

import "fmt"

type People struct {
	name string
	age  int
}

type Teacher struct {
	classroom string
	people    People
}

func main() {
	teacher := Teacher{"2006", People{"Aomsir", 20}}
	fmt.Println(teacher)
	fmt.Println(teacher.people.name, teacher.people)
}
