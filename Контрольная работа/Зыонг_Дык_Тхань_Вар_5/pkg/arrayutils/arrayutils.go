/*
Array utilities package
Description: Provides algorithms for array manipulations
Student: Duong_Duc_Thanh
License: GPLv3
**/

package arrayutils

// CircularShiftRight performs an in-place circular right shift of an integer array by n positions

func CircularShiftRight(arr []int, n int) {
	length := len(arr)

	// Handle edge cases
	if length <= 1 || n == 0 {
		return
	}

	// Normalize shift amount (in case n > length)
	n = n % length
	if n == 0 {
		return
	}

	// Reverse the entire array
	reverse(arr, 0, length-1)

	// Reverse the first n elements
	reverse(arr, 0, n-1)

	// Reverse the remaining elements
	reverse(arr, n, length-1)
}

// reverse reverses the elements in the array from start to end inclusive
func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
