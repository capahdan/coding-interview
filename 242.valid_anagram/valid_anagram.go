package valid_anagram

func ValidAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	arrS := []string{}
	arrT := []string{}

	for i := range s {
		arrS = append(arrS, string(s[i]))
	}

	for j := range t {
		arrT = append(arrT, string(t[j]))
	}

	arrS = mergeSort(arrS)
	arrT = mergeSort(arrT)

	for k := range arrS {
		if arrS[k] != arrT[k] {
			return false
		}
	}

	return true
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

func merge(a, b []string) []string {
	result := make([]string, len(a)+len(b)-1)
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}
