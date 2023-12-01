package group_anagram

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	var result [][]string

	anagramMap := make(map[string][]string)

	for _, str := range strs {
		splitted := strings.Split(str, "")
		sort.Strings(splitted)
		sortedStr := strings.Join(splitted, "")

		if data, ok := anagramMap[sortedStr]; ok {
			anagramMap[sortedStr] = append(data, str)
		} else {

			anagramMap[sortedStr] = []string{str}
		}
	}

	for _, value := range anagramMap {
		result = append(result, value)
	}

	return result
}

// kita harus tau dulu apakah ini adalah anagram atau tidak
// cara nya kita harus sorting isi nya dulu  bat -> abt, eat -> aet lalu
// lalu kalo sudah di sorting saattnya memasukan nilai awalnya kedalam array

//  for result

// eat -> aet
// tea -> aet
// tan -> ant
// ate -> aet
// nat -> ant
// bat -> abt
