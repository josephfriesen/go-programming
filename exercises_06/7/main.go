package main

import (
	"fmt"
)

func main() {
	// assign a func to a variable then call that func
	x := func() {
		fmt.Println("peepeepoopoo")
	}

	x()
}
