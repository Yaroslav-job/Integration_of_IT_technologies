package task2

import (
	"sync"
	"testing"
)

func TestGenerateAndSum(t *testing.T) {
	tests := []struct {
		name  string
		count int
	}{
		{
			name:  "Sum 0 numbers",
			count: 0,
		},
		{
			name:  "Sum 1 number",
			count: 1,
		},
		{
			name:  "Sum 10 numbers",
			count: 10,
		},
		{
			name:  "Sum 100 numbers",
			count: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum := GenerateAndSum(tt.count)

			// 1. If count = 0, sum = 0
			if tt.count == 0 && sum != 0 {
				t.Errorf("Expected sum of 0 for count 0, got %d", sum)
			}

			// 2. For positive counts, sum should be positive
			if tt.count > 0 && sum <= 0 {
				t.Errorf("Expected positive sum for count %d, got %d", tt.count, sum)
			}

			// 3. Sum should be within reasonable bounds
			if sum < tt.count || sum > tt.count*100 {
				t.Errorf("Sum %d is outside reasonable bounds for count %d", sum, tt.count)
			}
		})
	}
}

func TestCalculateSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{
			name:     "Empty list",
			numbers:  []int{},
			expected: 0,
		},
		{
			name:     "Single number",
			numbers:  []int{42},
			expected: 42,
		},
		{
			name:     "Multiple numbers",
			numbers:  []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "With negative numbers",
			numbers:  []int{-1, 2, -3, 4, -5},
			expected: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create channels
			in := make(chan int)
			out := make(chan int)

			// Start calculator goroutine
			go calculateSum(in, out)

			// Send numbers
			go func() {
				for _, num := range tt.numbers {
					in <- num
				}
				close(in)
			}()

			// Get result
			result := <-out

			// Check result
			if result != tt.expected {
				t.Errorf("calculateSum() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestNoRaceConditions verifies there are no race conditions by running multiple concurrent operations
func TestNoRaceConditions(t *testing.T) {
	const (
		numRoutines = 10
		iterations  = 10
	)

	var wg sync.WaitGroup
	wg.Add(numRoutines)

	// Run multiple GenerateAndSum operations concurrently
	for i := 0; i < numRoutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				GenerateAndSum(10)
			}
		}()
	}

	wg.Wait()

}
