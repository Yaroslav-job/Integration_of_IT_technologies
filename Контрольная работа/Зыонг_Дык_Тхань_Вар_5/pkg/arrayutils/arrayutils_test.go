package arrayutils

import (
	"reflect"
	"testing"
)

// TestCircularShiftRight tests the circular right shift algorithm
func TestCircularShiftRight(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		shift    int
		expected []int
	}{
		{
			name:     "Example from task",
			input:    []int{1, 2, 3, 4, 5},
			shift:    2,
			expected: []int{4, 5, 1, 2, 3},
		},
		{
			name:     "Shift by 0",
			input:    []int{1, 2, 3, 4, 5},
			shift:    0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Shift by array length",
			input:    []int{1, 2, 3, 4, 5},
			shift:    5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Shift by more than array length",
			input:    []int{1, 2, 3, 4, 5},
			shift:    7,
			expected: []int{4, 5, 1, 2, 3}, // Equivalent to shift by 2
		},
		{
			name:     "Empty array",
			input:    []int{},
			shift:    3,
			expected: []int{},
		},
		{
			name:     "Single element array",
			input:    []int{42},
			shift:    5,
			expected: []int{42},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create a copy to avoid modifying the original test data
			arrCopy := make([]int, len(test.input))
			copy(arrCopy, test.input)

			CircularShiftRight(arrCopy, test.shift)

			if !reflect.DeepEqual(arrCopy, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, arrCopy)
			}
		})
	}
}
