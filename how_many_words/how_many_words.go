package how_many_words

import (
	"sort"
	"strings"
)

// func CountSentences(wordSet []string) int {
// 	var result [][]string

// 	anagramMap := make(map[string][]string)

// 	for _, str := range strs {
// 		splitted := strings.Split(str, "")
// 		sort.Strings(splitted)
// 		sortedStr := strings.Join(splitted, "")

// 		if data, ok := anagramMap[sortedStr]; ok {
// 			anagramMap[sortedStr] = append(data, str)
// 		} else {

// 			anagramMap[sortedStr] = []string{str}
// 		}
// 	}

// 	for _, value := range anagramMap {
// 		result = append(result, value)
// 	}

// 	return result
// }

func CountSentences(wordSet []string, sentences []string) []int64 {
	anagramMap := make(map[string][]string)

	for _, str := range wordSet {
		splitted := strings.Split(str, "")
		sort.Strings(splitted)
		sortedStr := strings.Join(splitted, "")

		if data, ok := anagramMap[sortedStr]; ok {
			anagramMap[sortedStr] = append(data, str)
		} else {

			anagramMap[sortedStr] = []string{str}
		}
	}

	counts := make([]int64, len(sentences))

	for i, sentence := range sentences {
		// Split the sentence into words
		words := strings.Fields(sentence)

		// Initialize count for current sentence
		count := int64(1)

		for _, word := range words {
			splitted := strings.Split(word, "")
			sort.Strings(splitted)
			sortedStr := strings.Join(splitted, "")

			if data, ok := anagramMap[sortedStr]; ok && len(data) > 1 {
				count *= int64(len(data))
			}
		}
		counts[i] = count
	}

	return counts
}
