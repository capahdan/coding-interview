package palindrome

func Palindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	first := 0
	last := len(s) - 1

	for first < last {
		if s[first] != s[last] {
			return false
		}
		first++
		last--
	}
	return true
}
