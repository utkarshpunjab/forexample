package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("This is method speak attached to person using pointer receiver", p.name, p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := &person{
		name: "James Bond",
		age:  32,
	}
	p2 := person{
		name: "Jenny Ross",
		age:  27,
	}
	saySomething(p1)
	p2.speak()
	saySomething(&p2)
}
