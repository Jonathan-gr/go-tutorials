package main

import (
	"fmt"
	"sync"
)

func testingRaceConditions(wg *sync.WaitGroup) int32 {
	var state int32
	var mu sync.Mutex
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			mu.Lock()
			defer wg.Done()
			state += int32(i)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	return state
}

func main() {
	wg := sync.WaitGroup{}
	res := testingRaceConditions(&wg)

	fmt.Println(res) // deterministic result
}
