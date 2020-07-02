package main

import (
	"fmt"
)

func main() {
	// Create a func which returns a func
	// assign the returned func to a variable
	// call the returned func

	x := peepee()
	x()
}

func peepee() func() {
	return func() {
		fmt.Println("poopoo")
	}
}
