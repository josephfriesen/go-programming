package main

import (
	"fmt"
)

func main() {
	// Closure is when we have “enclosed” the scope of a variable in some code block.
	// For this hands-on exercise, create a func which “encloses” the scope of a variable:

	peepee := func(n int) int {
		x := 400
		return n + x
	}

	fmt.Println(peepee(33)) // 33
	// fmt.Println(x) // ./main.go:17:14: undefined: x
}
