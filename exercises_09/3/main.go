package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	steps := 100
	count := 0

	wg.Add(steps)

	for i := 1; i <= steps; i++ {
		go func() {
			current := count
			runtime.Gosched()
			current++
			count = current
			fmt.Println("STEP", i, "\tCOUNT", count)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("All done, final count:", count)
}
