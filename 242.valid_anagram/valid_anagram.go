package valid_anagram

import "strings"

func ValidAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	splitedS := strings.Split(s, "")
	splittedT := strings.Split(t, "")

	sortedS := mergeSort(splitedS)
	sortedT := mergeSort(splittedT)

	for i := range sortedS {
		if sortedS[i] != sortedT[i] {
			return false
		}
	}
	return true
}

func merge(left, right []string) []string {
	result := make([]string, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func mergeSort(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}
