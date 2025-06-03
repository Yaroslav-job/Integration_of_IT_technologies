# Go Algorithms and Concurrency

This project implements common array algorithms and concurrent programming patterns in Go.

## Features

1. **Task 1: Array Circular Shift Algorithm**
   - Implementation of an in-place circular right shift algorithm
   - Time complexity: O(n)
   - Space complexity: O(1)
   - Example: shifts [1, 2, 3, 4, 5] by 2 positions to get [4, 5, 1, 2, 3]

2. **Task 2: Concurrent Number Generation and Summation**
   - Uses goroutines to generate random numbers
   - Sends the numbers through channels
   - Calculates the sum in a separate goroutine
   - Demonstrates proper coordination and channel closing

### How to Use
# Build the project
// make build
# Run the application
// make run
# Run all tests
// make test
