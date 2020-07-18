package main

import (
	"fmt"
	"math/rand"
	"time"
)

type number struct {
	n           int
	goroutineID int
}

func main() {
	fmt.Println("Generate 100 random integers from 0 and 100")
	rand.Seed(time.Now().UnixNano())

	c := gen()

	fmt.Println("n\tgoroutine")
	fmt.Println("=================================")
	for j := 0; j < 100; j++ {
		v := <-c
		fmt.Println(v.n, "\t", v.goroutineID)
	}

	fmt.Println("Done.")
}

func gen() <-chan number {
	c := make(chan number)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- number{
					n:           rand.Intn(100),
					goroutineID: i,
				}
				time.Sleep(250 * time.Millisecond)
			}
		}()
	}
	return c
}
