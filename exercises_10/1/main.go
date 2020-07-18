package main

import (
	"fmt"
)

// 1
// func main() {
// 	c := make(chan int)

// 	go func() {
// 		c <- 42
// 	}()

// 	fmt.Println(<-c)
// }

// 2
func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
