package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// (1) append 52, print
	x = append(x, 52)
	fmt.Println(x) // [42 43 44 45 46 47 48 49 50 51 52]

	// (2) in one statement, append 53, 54, 55
	x = append(x, 53, 54, 55)
	fmt.Println(x) // [42 43 44 45 46 47 48 49 50 51 52 53 54 55]

	// (3) append [56, 57, 58, 59, 60]
	x = append(x, []int{56, 57, 58, 59, 60}...)
	fmt.Println(x) // [42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60]
}
