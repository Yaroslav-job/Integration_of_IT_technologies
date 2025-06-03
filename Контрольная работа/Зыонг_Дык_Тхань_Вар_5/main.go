/*
File: main.go
Description: Entry point for the Go algorithms application
Student: Duong_Duc_Thanh
License: GPLv3
**/

package main

import (
	"fmt"

	"github.com/user/goalgorithms/pkg/task1"
	"github.com/user/goalgorithms/pkg/task2"
)

func main() {
	// Task 1
	fmt.Println("=== Task 1 ===")
	arr := []int{1, 2, 3, 4, 5}
	n := 2
	fmt.Printf("Original array: %v\n", arr)
	fmt.Printf("Shifting right by %d positions\n", n)

	shifted := make([]int, len(arr))
	copy(shifted, arr)
	task1.CircularShiftRight(shifted, n)

	fmt.Printf("Result: %v\n", shifted)

	// Task 2
	fmt.Println("\n=== Task 2 ===")
	fmt.Println("Generating 100 random numbers and calculating their sum concurrently")

	sum := task2.GenerateAndSum(100)
	fmt.Printf("Sum of generated numbers: %d\n", sum)
}
