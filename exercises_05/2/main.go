package main

import (
	"fmt"
)

func main() {
	// Take the code from the previous exercise, then store the values of type person in a map
	// with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

	type person struct {
		firstName   string
		lastName    string
		favIceCream []string
	}

	p1 := person{
		firstName:   "groucho",
		lastName:    "marx",
		favIceCream: []string{"chocolate peanut butter", "chocolate"},
	}
	p2 := person{
		firstName:   "harpo",
		lastName:    "marx",
		favIceCream: []string{"rocky road", "vanilla"},
	}

	people := map[string]person{
		"groucho": p1,
		"harpo":   p2,
	}

	fmt.Println("Person 1")
	fmt.Printf("\t%v ", people["groucho"].firstName)
	fmt.Printf("%v\n", people["groucho"].lastName)
	fmt.Printf("\tFavorite ice cream flavors:\n")
	for _, v := range people["groucho"].favIceCream {
		fmt.Println("\t\t", v)
	}
	fmt.Println("Person 2")
	fmt.Printf("\t%v ", people["harpo"].firstName)
	fmt.Printf("%v\n", people["harpo"].lastName)
	fmt.Printf("\tFavorite ice cream flavors:\n")
	for _, v := range people["harpo"].favIceCream {
		fmt.Println("\t\t", v)
	}

	// output
	// Person 1
	//     groucho marx
	//     Favorite ice cream flavors:
	//              chocolate peanut butter
	//              chocolate
	// Person 2
	//     harpo marx
	//     Favorite ice cream flavors:
	//              rocky road
	//              vanilla
}
