package main

import (
	"fmt"
	"palindrome-checker/internal/palindrome"
)

func main() {
	testStrings := []string{"level", "hello", "madam", "racecar", "world"}

	for _, str := range testStrings {
		if palindrome.IsPalindrome(str) {
			fmt.Printf("\"%s\" is a palindrome\n", str)
		} else {
			fmt.Printf("\"%s\" is NOT a palindrome\n", str)
		}
	}
}
