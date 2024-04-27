package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1) // Want the application to use to 1 CPU core
	s := time.Now()

	start := 1
	end := int(1e7)

	// Calculate the number of goroutines needed
	numGoroutines := runtime.NumCPU()
	chunkSize := (end - start + 1) / numGoroutines

	// Create a channel to collect the results
	results := make(chan int)

	// start goroutines
	for i := 0; i < numGoroutines; i++ {
		start := start + i*chunkSize
		end := start + chunkSize - 1
		if i == numGoroutines-1 {
			end = int(1e7)
		}
		go sum(start, end, results)
	}

	totalSum := 0
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-results
	}
	fmt.Println(totalSum)
	fmt.Println(time.Since(s))
}

func sum(start, end int, results chan<- int) {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	results <- sum
}
