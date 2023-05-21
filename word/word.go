package word

import "strings"

func WordCount(s string) int {
	if strings.Count(s, " ") == 0 {
		return 1
	}

	return strings.Count(s, " ") + 1
}
