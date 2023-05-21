package word

import "strings"

func WordCount(s string) int {
	return strings.Count(s, " ")
}
