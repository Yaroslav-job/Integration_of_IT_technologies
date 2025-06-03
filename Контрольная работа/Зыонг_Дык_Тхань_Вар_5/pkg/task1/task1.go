/*
File: task1.go
Description: Implementation of array algorithms including circular shift
Student: Duong_Duc_Thanh
License: GPLv3
**/

package task1

// CircularShiftRight performs an in-place circular shift of an integer array by n positions to the right. This implementation does not use any additional array.
// Time complexity: O(n), where n is the length of the array
// Space complexity: O(1)

func CircularShiftRight(arr []int, shiftBy int) {
	if len(arr) == 0 || shiftBy <= 0 {
		return
	}

	// Normalize shift in case it's larger than array length
	shiftBy = shiftBy % len(arr)
	if shiftBy == 0 {
		return
	}

	// Algorithm:
	// 1: Reverse the entire array
	reverse(arr, 0, len(arr)-1)

	// 2: Reverse first 'shiftBy' elements
	reverse(arr, 0, shiftBy-1)

	// 3: Reverse remaining elements
	reverse(arr, shiftBy, len(arr)-1)
}

// reverse a slice in-place from start to end indices inclusive
func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
