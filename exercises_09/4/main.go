package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	steps := 100
	count := 0
	var mu sync.Mutex

	wg.Add(steps)

	for i := 1; i <= steps; i++ {
		go func() {
			mu.Lock()
			current := count
			current++
			count = current
			fmt.Println("\tCOUNT", count)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("All done, final count:", count)
}
