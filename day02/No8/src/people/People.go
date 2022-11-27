package people

type People struct {
	name string
	age  int
}

func (p *People) SetName(name string) {
	p.name = name
}

func (p *People) SetAge(age int) {
	p.age = age
}

func (p People) GetName() string {
	return p.name
}

func (p People) GetAge() int {
	return p.age
}
