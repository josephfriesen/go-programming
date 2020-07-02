package main

import (
	"fmt"
)

func main() {
	// Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
	// 	first name
	// 	last name
	// 	favorite ice cream flavors
	// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
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

	fmt.Println("Person 1")
	fmt.Printf("\t%v ", p1.firstName)
	fmt.Printf("%v\n", p1.lastName)
	fmt.Printf("\tFavorite ice cream flavors:\n")
	for _, v := range p1.favIceCream {
		fmt.Println("\t\t", v)
	}
	fmt.Println("Person 2")
	fmt.Printf("\t%v ", p2.firstName)
	fmt.Printf("%v\n", p2.lastName)
	fmt.Printf("\tFavorite ice cream flavors:\n")
	for _, v := range p2.favIceCream {
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
