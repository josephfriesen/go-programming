package main

import (
	"fmt"
)

func main() {
	// (1) Create a new type: vehicle.
	// 	The underlying type is a struct.
	// 	The fields:
	// 		doors
	// 		color
	// (2) Create two new types: truck & sedan.
	// 		The underlying type of each of these new types is a struct.
	// 		Embed the “vehicle” type in both truck & sedan.
	// 		Give truck the field “fourWheel” which will be set to bool.
	// 		Give sedan the field “luxury” which will be set to bool.
	// (3) Using the vehicle, truck, and sedan structs:
	// 		using a composite literal, create a value of type truck and assign values to the fields;
	// 		using a composite literal, create a value of type sedan and assign values to the fields.
	// (4) Print out each of these values.
	// (5) Print out a single field from each of these values.

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	truckasaurus := truck{
		vehicle: vehicle{
			doors: 2,
			color: "green",
		},
		fourWheel: true,
	}
	sux6000 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: false,
	}

	fmt.Println(truckasaurus)                        // {{2 green} true}
	fmt.Printf("\tDoors:\t%v\n", truckasaurus.doors) // Doors:  2
	fmt.Println(sux6000)                             // {{4 blue} false}
	fmt.Printf("\tColor:\t%v\n", sux6000.color)      // Color:  blue

}
