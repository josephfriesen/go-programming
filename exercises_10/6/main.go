package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Generate 100 random integers from 0 and 100")
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))

	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- rand.Intn(100)
			time.Sleep(5 * time.Millisecond)
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Done.")
}
