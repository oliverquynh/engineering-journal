package main

import (
	"fmt"
	"sync"
)

func main() {
	// Shared resource: The product stock counter
	counter := 0

	// Configuration for the concurrency test
	totalThreads := 5
	iterations := 100000

	// WaitGroup is used to wait for all goroutines to finish
	var wg sync.WaitGroup

	fmt.Printf("Starting %d concurrent threads, each incrementing %d times...\n", totalThreads, iterations)
	fmt.Printf("Expected Theoretical Result: %d\n", totalThreads*iterations)
	fmt.Println("--------------------------------------------------")

	// Spawning multiple concurrent threads (goroutines)
	for i := 0; i < totalThreads; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				// Race Condition happens here: multiple threads
				// accessing and modifying the same memory address concurrently
				counter++
			}
		}()
	}

	// Block until all threads complete their work
	wg.Wait()

	// Output the final actual result stored in RAM
	fmt.Printf("ACTUAL RESULT IN RAM: %d\n", counter)

	if counter != totalThreads*iterations {
		loss := (totalThreads * iterations) - counter
		fmt.Printf("⚠️ DATA CORRUPTION DETECTED! Lost: %d units.\n", loss)
	} else {
		fmt.Println("🎉 Clean run! No race condition caught this time.")
	}
}


