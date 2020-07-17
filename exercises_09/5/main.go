package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	steps := 100
	var count int64

	wg.Add(steps)

	for i := 1; i <= steps; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			fmt.Println("\tCOUNT", atomic.LoadInt64(&count))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("All done, final count:", count)
}
