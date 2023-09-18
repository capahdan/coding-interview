package length_of_last_word

import "strings"

func LenghtOfLastWord(s string) int {

	trimmed := trim(s, " ")

	splitted := strings.Split(trimmed, " ")

	return len(splitted[len(splitted)-1])
}

func trim(input, trim string) string {
	temp := input
	for i := range input {
		if string(input[i]) == trim {
			temp = temp[i+1:]
		}
	}
	for i := len(input) - 1; i >= 0; i++ {
		if string(input[i]) == trim {
			temp = temp[:i-1]
		}
	}
	return temp
}
