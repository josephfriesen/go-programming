package main

import (
	"fmt"
)

func main() {
	// create a func with the identifier foo that returns an int
	// create a func with the identifier bar that returns an int and a string
	// call both funcs
	// print out their results
	x := peepee()
	y, s := poopoo()
	fmt.Println(x)    // 2
	fmt.Println(y, s) // 3 peepeepoopoo
}

func peepee() int {
	return 2
}
func poopoo() (int, string) {
	return 3, "peepeepoopoo"
}
