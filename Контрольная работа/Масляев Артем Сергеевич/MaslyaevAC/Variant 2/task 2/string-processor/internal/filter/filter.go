// Package filter фильтрует строки, содержащие символ 'a'.
package filter

import (
	"fmt"
)

// FilterStrings читает строки из канала и фильтрует их по наличию 'a'.
func FilterStrings(ch <-chan string) {
	for str := range ch {
		if containsA(str) {
			fmt.Println("Contains 'a':", str)
		}
	}
}

// containsA проверяет, есть ли в строке символ 'a'.
func containsA(s string) bool {
	for _, r := range s {
		if r == 'a' {
			return true
		}
	}
	return false
}
