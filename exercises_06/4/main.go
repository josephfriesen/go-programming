package main

import (
	"fmt"
)

func main() {
	// Create a user defined struct with
	// 	the identifier “person”
	// 	the fields:
	// 		first
	// 		last
	// 		age
	// attach a method to type person with
	// 	the identifier “speak”
	// 	the method should have the person say their name and age
	// create a value of type person
	// 	call the method from the value of type person

	groucho := person{
		first: "Groucho",
		last:  "Marx",
		age:   129,
	}

	groucho.speak() // hello my name is Groucho Marx and i am 129 years old
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("hello my name is", p.first, p.last, "and i am", p.age, "years old")
}
