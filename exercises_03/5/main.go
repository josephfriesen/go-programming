package main

import (
	"fmt"
)

func main() {
	// print i mod 4 for i = 10, 11, ... , 100
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v mod 4 = \t %v\n", i, i%4)
	}
}
