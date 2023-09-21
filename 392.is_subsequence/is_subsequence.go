package valid_palindrome

func IsSubsequence(s, t string) bool {
	sIdx := 0
	for i := 0; i < len(t) && sIdx < len(s); i++ {
		if t[i] == s[sIdx] {
			sIdx++
		}
	}
	return sIdx == len(s)
}
