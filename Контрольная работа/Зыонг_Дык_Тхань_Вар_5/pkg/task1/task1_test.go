package task1

import (
	"reflect"
	"testing"
)

func TestCircularShiftRight(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		shiftBy  int
		expected []int
	}{
		{
			name:     "Empty array",
			arr:      []int{},
			shiftBy:  3,
			expected: []int{},
		},
		{
			name:     "Zero shift",
			arr:      []int{1, 2, 3, 4, 5},
			shiftBy:  0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Shift by 2 positions",
			arr:      []int{1, 2, 3, 4, 5},
			shiftBy:  2,
			expected: []int{4, 5, 1, 2, 3},
		},
		{
			name:     "Shift by array length",
			arr:      []int{1, 2, 3, 4, 5},
			shiftBy:  5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Shift by more than array length",
			arr:      []int{1, 2, 3, 4, 5},
			shiftBy:  7,
			expected: []int{4, 5, 1, 2, 3},
		},
		{
			name:     "Negative shift (should be ignored)",
			arr:      []int{1, 2, 3, 4, 5},
			shiftBy:  -2,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Single element",
			arr:      []int{42},
			shiftBy:  10,
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CircularShiftRight(tt.arr, tt.shiftBy)
			if !reflect.DeepEqual(tt.arr, tt.expected) {
				t.Errorf("CircularShiftRight() = %v, want %v", tt.arr, tt.expected)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		start    int
		end      int
		expected []int
	}{
		{
			name:     "Reverse entire array",
			arr:      []int{1, 2, 3, 4, 5},
			start:    0,
			end:      4,
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Reverse first half",
			arr:      []int{1, 2, 3, 4, 5},
			start:    0,
			end:      2,
			expected: []int{3, 2, 1, 4, 5},
		},
		{
			name:     "Reverse second half",
			arr:      []int{1, 2, 3, 4, 5},
			start:    2,
			end:      4,
			expected: []int{1, 2, 5, 4, 3},
		},
		{
			name:     "Reverse single element (no change)",
			arr:      []int{1, 2, 3, 4, 5},
			start:    2,
			end:      2,
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.arr, tt.start, tt.end)
			if !reflect.DeepEqual(tt.arr, tt.expected) {
				t.Errorf("reverse() = %v, want %v", tt.arr, tt.expected)
			}
		})
	}
}
