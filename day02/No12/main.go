package main

import "fmt"

type Live interface {
	run()
}

type People struct{}

type Animal struct{}

func (people People) run() {
	fmt.Println("人正在跑步")
}

func (animal Animal) run() {
	fmt.Println("动物正在奔跑")
}

func allrun(live Live) {
	live.run()
}

func main() {
	people := &People{}
	animal := &Animal{}

	allrun(people)
	allrun(animal)
}
