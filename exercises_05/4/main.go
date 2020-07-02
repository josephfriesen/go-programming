package main

import (
	"fmt"
)

func main() {
	// Create and use an anonymous struct
	truckasaurus := struct {
		doors     int
		color     string
		fourWheel bool
	}{
		doors:     2,
		color:     "green",
		fourWheel: true,
	}

	fmt.Println(truckasaurus) // {2 green true}
}
