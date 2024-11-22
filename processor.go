package main

import (
	"sync"
)

func processAndGather(in <-chan int, processor func(int) int, concurrentNumber int) []int {
	out := make(chan int, concurrentNumber)

	var wg sync.WaitGroup
	wg.Add(concurrentNumber)

	for i := 0; i < concurrentNumber; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				if v < 0 {
					break
				}
				out <- processor(v)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var result []int
	for v := range out {
		result = append(result, v)
	}

	return result
}
