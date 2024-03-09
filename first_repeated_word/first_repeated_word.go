package first_repeated_word

import "strings"

func FirstRepeatedWord(sentence string) string {
	mapStr := make(map[string]bool)

	sentence = replaceSpecialChars(sentence)

	words := strings.Fields(sentence)
	for _, word := range words {
		if _, ok := mapStr[word]; ok {
			return word
		}
		mapStr[word] = true
	}

	return ""
}

func replaceSpecialChars(sentence string) string {
	// Define a map of special characters to be replaced with space
	specialChars := ",.-:"

	// Replace special characters with spaces
	for _, char := range specialChars {
		sentence = strings.ReplaceAll(sentence, string(char), " ")
	}

	return sentence
}
