package main

import "fmt"

type People struct {
	Name   string
	Weight int
}

func (p *People) run() {
	fmt.Println(p.Name, "正在跑步,当前体重为", p.Weight)
	p.Weight--
}

func main() {
	people := People{"Aomsir", 56}
	for i := 0; i < 10; i++ {
		people.run()
	}

}
