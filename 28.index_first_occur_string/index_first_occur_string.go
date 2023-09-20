package index_first_occur_string

import "strings"

func StrStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	if !strings.Contains(haystack, needle) {
		return -1
	}

	return strings.Index(haystack, needle)
}
