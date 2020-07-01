package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is the most awesome class in the entire world for having fun and learning the GO programming language.")

	peepee()

	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	poopoo()
}

func peepee() {
	fmt.Println("I'm in foo.")
}

func poopoo() {
	fmt.Println("and then we exited")
}

// control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional
