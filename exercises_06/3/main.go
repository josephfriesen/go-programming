package main

import (
	"fmt"
)

func main() {
	// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
	peepee("start")
	defer peepee("i ran after the rest of the func main() code block completed")
	peepee("...wait for it")
	peepee("...still waiting")
	peepee("...almost there")

	// output:
	// start
	// ...wait for it
	// ...still waiting
	// ...almost there
	// i ran after the rest of the func main() code block completed
}

func peepee(s string) {
	fmt.Println(s)
}
