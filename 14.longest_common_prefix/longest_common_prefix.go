package longest_common_prefix

import (
	"math"
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

	minLen := int(math.Min(float64(len(first)), float64(len(last))))

	for i := 0; i < minLen; i++ {
		if first[i] != last[i] {
			return resultStr
		}
		resultStr += string(first[i])
	}

	return resultStr
}
