package valid_palindrome

import (
	"strings"
)

// func ValidPalindrome(s string) bool {

// 	if len(s) <= 1 {
// 		return true
// 	}

// 	loweredStr := strings.ToLower(s)

// 	pattern := regexp.MustCompile("[[:alnum:]]+")
// 	matches := pattern.FindAllString(loweredStr, -1)

// 	filteredStr := ""
// 	for _, match := range matches {
// 		filteredStr += match
// 	}

// 	i := 0
// 	j := len(filteredStr) - 1
// 	for i <= j {
// 		if filteredStr[i] != filteredStr[j] {
// 			return false
// 		}
// 		i++
// 		j--
// 	}
// 	return true
// }

func ValidPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func isAlphanumeric(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
}
