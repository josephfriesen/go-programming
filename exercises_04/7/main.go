package main

import "fmt"

func main() {
	// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
	//		"James", "Bond", "Shaken, not stirred"
	//		"Miss", "Moneypenny", "Helloooooo, James."
	// Range over the records, then range over the data in each record.

	x := []string{"James", "Bond", "Shaken, not stirred"} // [James Bond Shaken, not stirred]
	y := []string{"Miss", "Moneypenny", "Hello, James"}   // [Miss Moneypenny Hello, James]

	fmt.Println(x)
	fmt.Println(y)

	z := [][]string{x, y}
	fmt.Println(z)
	for _, record := range z {
		fmt.Println(record)
		for _, element := range record {
			fmt.Println(element)
		}
	}
}
