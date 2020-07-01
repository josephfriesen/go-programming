package main

import (
	"fmt"
)

func main() {
	// print ASCII code for each uppercase letter of the alphabet 3 times
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
		fmt.Printf("\n")
	}
}
