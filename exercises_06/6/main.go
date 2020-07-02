package main

import (
	"fmt"
)

func main() {
	// build and use an anonymous func
	func() {
		fmt.Println("peepeepoopoo")
	}()
}
