package main

import (
	"fmt"
)

func main() {
	// (1) create new slices from the following slice
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(slice[:5])  // [42, 43, 44, 45, 46]
	fmt.Println(slice[5:])  // [47, 48, 49, 50, 51]
	fmt.Println(slice[2:7]) // [44, 45, 46, 47, 48]
	fmt.Println(slice[1:6]) // [43, 44, 45, 46, 47]
}
