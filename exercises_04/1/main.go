package main

import (
	"fmt"
)

func main() {
	// (1) use a composite literal to create an array which holds 5 values of type int, assign values to each position
	arr := [5]int{10, 21, 32, 43, 54}
	fmt.Println(arr)

	// (2) range over array and print values
	for _, v := range arr {
		fmt.Println(v)
	}

	// (3) print the type of the array using printf
	fmt.Printf("%T\n", arr)
}
