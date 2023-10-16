package isomorphic_string

func IsomorphicString(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Create two maps for s & t strings...
	map1 := make(map[byte]int)
	map2 := make(map[byte]int)

	// Traverse all elements through the loop...
	for idx := 0; idx < len(s); idx++ {
		// Compare the maps, if not equal, return false...

		if map1[s[idx]] != map2[t[idx]] {
			return false
		}

		// Insert each character from string s and t into separate map...
		map1[s[idx]] = idx + 1
		map2[t[idx]] = idx + 1
	}

	return true
}
