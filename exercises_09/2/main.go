package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hi my name is", p.first, p.last)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	dude := person{
		first: "Adam",
		last:  "West",
	}

	// saySomething(dude)
	// OUTPUT: ./main.go:30:14: cannot use dude (type person) as type human in argument to saySomething:
	// 		person does not implement human (speak method has pointer receiver)

	saySomething(&dude)
	// Hi my name is Adam West
}
