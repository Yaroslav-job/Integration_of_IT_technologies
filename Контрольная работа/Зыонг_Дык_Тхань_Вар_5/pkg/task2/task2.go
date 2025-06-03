/*
File: task2.go
Description: Implementation of concurrent programming patterns using goroutines and channels
Student: Duong_Duc_Thanh
License: GPLv3
*/

package task2

import (
	"math/rand"
	"sync"
	"time"
)

// GenerateAndSum creates multiple goroutines that generate random numbers. Another goroutine calculates the sum of these numbers.
// Returns the final sum after all numbers have been generated and processed.
func GenerateAndSum(count int) int {
	// Create a channel for number communication
	numChan := make(chan int)

	// Number of generator goroutines
	const numGenerators = 5

	// Use WaitGroup to ensure all generators have finished
	var wg sync.WaitGroup
	wg.Add(numGenerators)

	// Initialize random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Start the calculator goroutine
	sumChan := make(chan int)
	go calculateSum(numChan, sumChan)

	// Start generator goroutines
	numbersPerGenerator := count / numGenerators
	remainingNumbers := count % numGenerators

	for i := 0; i < numGenerators; i++ {
		// Distribute the count evenly among generators
		genCount := numbersPerGenerator
		if i < remainingNumbers {
			genCount++
		}

		go generateNumbers(numChan, genCount, r.Int63(), &wg)
	}

	// Wait for all generators to finish, then close the channel
	go func() {
		wg.Wait()
		close(numChan)
	}()

	return <-sumChan
}

// generateNumbers sends 'count' random numbers to the provided channel.

func generateNumbers(ch chan<- int, count int, seed int64, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create a local random source for thread safety
	localRand := rand.New(rand.NewSource(seed))

	for i := 0; i < count; i++ {
		// Generate random number between 1 and 100
		num := localRand.Intn(100) + 1
		ch <- num

		time.Sleep(time.Millisecond)
	}
}

// calculateSum receives numbers from the input channel, calculates their sum, and sends the final result.
func calculateSum(in <-chan int, out chan<- int) {
	sum := 0

	for num := range in {
		sum += num
	}

	out <- sum
	close(out)
}
