package longest_substring

//  s = "abcabcbb"

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func LengthOfLongestSubstring(s string) int {
	left := 0
	set := make(map[byte]bool)
	maxSize := 0

	for i := 0; i < len(s); i++ {
		for set[s[i]] {
			delete(set, s[left])
			left++
		}
		set[s[i]] = true
		maxSize = max(maxSize, i-left+1)
	}

	return maxSize
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
