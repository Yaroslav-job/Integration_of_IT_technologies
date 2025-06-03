// Package palindrome предоставляет функцию для проверки строки на палиндром.
package palindrome

// IsPalindrome проверяет, является ли строка палиндромом без использования доп. структур данных.
func IsPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
