package main

import (
	"fmt"
)

func main() {
	// create a func with the identifier foo that
	//		takes in a variadic parameter of type int
	//		pass in a value of type []int into your func (unfurl the []int)
	// 		returns the sum of all values of type int passed in
	// create a func with the identifier bar that
	// 		takes in a parameter of type []int
	// 		returns the sum of all values of type int passed in

	fmt.Println(peepee(1, 2, 3, 4, 5))        // 15
	fmt.Println(poopoo([]int{1, 2, 3, 4, 5})) // 15
}

func peepee(n ...int) int {
	return poopoo(n)
}

func poopoo(arr []int) int {
	var total int
	for _, v := range arr {
		total += v
	}
	return total
}
