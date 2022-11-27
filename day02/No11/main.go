package main

import "fmt"

type Live interface {
	run(length int)
	eat()
}

type People struct {
	name string
}

// People类实现Live这个interface的run方法
func (people *People) run(length int) {
	fmt.Println(people.name, "正在跑步,跑了", length, "米")
}

// People类实现Live这个interface的run方法
func (people *People) eat() {
	fmt.Println(people.name, "正在吃饭")
}

func main() {
	people := People{"AOmsir"}
	people.run(100)
	people.eat()
}
