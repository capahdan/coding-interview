package longest_substring

//  s = "abcabcbb"

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// "babad"
// "cbbd"
func LongestPalindrom(s string) string {
	if len(s) == 1 {
		return s
	}
	str := ""
	if len(s) >= 2 {
		temp := s
		if isPalindrom(temp) {
			str = temp
		}
	}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			new := s[i:j]
			if isPalindrom(s[i:j]) {
				if len(new) > len(str) {
					str = new
				}
			}
		}
	}

	return str
}

func isPalindrom(s string) bool {
	start := 0
	lenght := len(s) - 1

	for start <= lenght {
		if s[start] != s[lenght] {
			return false
		}
		start++
		lenght--
	}
	return true
}
