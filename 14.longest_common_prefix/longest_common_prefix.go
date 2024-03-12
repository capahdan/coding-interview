package longest_common_prefix

import (
	"sort"
)

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Sort the input strings
	sort.Strings(strs)

	first := strs[0]
	last := strs[len(strs)-1]
	resultStr := ""

	minLen := min(len(first), len(last))

	for i := 0; i < minLen; i++ {
		if first[i] != last[i] {
			return resultStr
		}
		resultStr += string(first[i])
	}

	return resultStr
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
