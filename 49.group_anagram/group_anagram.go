package group_anagram

func GroupAnagrams(strs []string) [][]string {
	var result [][]string

	anagramMap := make(map[[26]int][]string)

	for _, str := range strs {
		k := [26]int{}
		for i := 0; i < len(str); i++ {
			k[str[i]-'a'] += 1
		}
		anagramMap[k] = append(anagramMap[k], str)
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
