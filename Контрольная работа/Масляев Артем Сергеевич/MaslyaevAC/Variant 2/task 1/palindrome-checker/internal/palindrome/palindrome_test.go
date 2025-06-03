package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"level", true},
		{"hello", false},
		{"madam", true},
		{"racecar", true},
		{"world", false},
		{"", true},         // пустая строка - палиндром
		{"a", true},        // одиночный символ - палиндром
		{"ab", false},      // два разных символа - не палиндром
		{"aba", true},      // классический палиндром
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := IsPalindrome(test.input)
			if result != test.expected {
				t.Errorf("For input %q expected %v but got %v", test.input, test.expected, result)
			}
		})
	}
}
