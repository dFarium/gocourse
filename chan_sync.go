package main

import (
	"time"
)

// func main() {
// 	done := make(chan int)

// 	go func() {
// 		// Simulate some work with a goroutine
// 		// After work is done, close the channel to signal completion
// 		fmt.Println("Goroutine is doing some work...")
// 		time.Sleep(2 * time.Second)
// 		done <- 47
// 	}()

// 	// Wait for the goroutine to signal completion
// 	result := <-done
// 	fmt.Println("Goroutine has completed with result:", result)
// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		ch <- 47
// 		time.Sleep(1 * time.Second)
// 		fmt.Println("Sent 47 to channel")
// 	}()

// 	// Wait for the goroutine to signal completion
// 	value := <-ch
// 	fmt.Println("Goroutine has completed with result:", value)
// }

func main() {
	numGoroutines := 3
	ch := make(chan int, numGoroutines)

	for i := range numGoroutines {
		go func(i int) {
			ch <- i
			time.Sleep(1 * time.Second)
		}(i)
	}

}
