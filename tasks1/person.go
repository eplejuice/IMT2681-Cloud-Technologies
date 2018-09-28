package main

type person struct {
	age int
}

func (p *person) addten() {
	p.age += 10
}

func add10(p *person) {
	p.age += 10
}
