package main

import (
	"fmt"
)

func main() {
	// pass a func into a func as an argument
	s := peepee(poopoo)
	fmt.Println(s) // peepee poopoo
}

func peepee(cb func() string) string {
	suffix := cb()
	return fmt.Sprint("peepee", " ", suffix)
}

func poopoo() string {
	return "poopoo"
}