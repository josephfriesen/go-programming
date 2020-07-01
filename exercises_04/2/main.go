package main

import (
	"fmt"
)

func main() {
	// (1) create a slice of type int assign 10 values
	slice := []int{12, 43, 12, 54, 33, 11, 66, 11, 12, 87}
	fmt.Println(slice)

	// (2) range over the slice and print the values out
	for _, v := range slice {
		fmt.Println(v)
	}

	// (3) print out the type of the slice
	fmt.Printf("%T\n", slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
