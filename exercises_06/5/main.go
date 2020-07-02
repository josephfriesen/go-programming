package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type square struct {
	sideLength float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (s square) area() float64 {
	return math.Pow(s.sideLength, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	// create a type SQUARE
	// create a type CIRCLE
	// attach a method to each that calculates AREA and returns it
	// circle area= π r 2
	// square area = L * W
	// create a type SHAPE that defines an interface as anything that has the AREA method
	// create a func INFO which takes type shape and then prints the area
	// create a value of type square
	// create a value of type circle
	// use func info to print the area of square
	// use func info to print the area of circle

	circ := circle{
		// radius: 25.0,
		radius: 12.345,
	}

	squ := square{
		// sideLength: 10.1,
		sideLength: 15,
	}

	info(circ) // 1963.4954084936207
	info(squ)  // 102.00999999999999
}
