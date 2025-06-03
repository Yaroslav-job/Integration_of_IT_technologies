/*
Project: Контрольная работа. Алгоритмы и Go Routines.
Description: Unit test для task1.
Author: Никита Мякишев (вариант 3)
License: GPLv3
History:
  - [03.04.2025 17:25]: Initial task1_test
*/

package task1

import (
	"cr1/src/task1"
	"testing"
)

// Unit-тест для проверки longestCommonPrefix
func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interspecies", "interstellar", "interstate"}, "inters"},
		{[]string{"a"}, "a"},
		{[]string{"abc", "abc", "abc"}, "abc"},
		{[]string{"", "b"}, ""},
	}

	for _, testCase := range testCases {
		result := task1.LongestCommonPrefix(testCase.input)
		if result != testCase.expected {
			t.Errorf("Для входных данных %v ожидалось %s, но получено %s", testCase.input, testCase.expected, result)
		}
	}
}
