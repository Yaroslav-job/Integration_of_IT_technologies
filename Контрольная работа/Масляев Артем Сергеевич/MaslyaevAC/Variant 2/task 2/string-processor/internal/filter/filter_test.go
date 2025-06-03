package filter

import "testing"

func TestContainsA(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"apple", true},
		{"hello", false},
		{"banana", true},
		{"xyz", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := containsA(test.input)
			if result != test.expected {
				t.Errorf("For input %q expected %v but got %v", test.input, test.expected, result)
			}
		})
	}
}
