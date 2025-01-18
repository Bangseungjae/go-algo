package _25_valid_palindrome

import (
	"unicode"
)

func isPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		switch {
		case !unicode.IsLetter(rune(s[start])) && !unicode.IsNumber(rune(s[start])):
			start++
			continue
		case !unicode.IsLetter(rune(s[end])) && !unicode.IsNumber(rune(s[end])):
			end--
			continue
		default:
			if unicode.ToLower(rune(s[start])) != unicode.ToLower(rune(s[end])) {
				return false
			}
			start++
			end--
		}
	}
	return true
}
