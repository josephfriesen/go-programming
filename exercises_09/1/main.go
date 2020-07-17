package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Num CPU\t", runtime.NumCPU())
	fmt.Println("Num goroutines:\t", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Hi i am the first line of the main goroutine")

	go func() {
		fmt.Println("Hi i am in the second goroutine")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hi i am in the third goroutine")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Num CPU\t", runtime.NumCPU())
	fmt.Println("Num goroutines:\t", runtime.NumGoroutine())
}
