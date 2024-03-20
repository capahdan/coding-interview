package permutation

func CheckInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) || len(s2) == 0 {
		return false
	}

	if len(s1) == 0 {
		return true
	}

	arr1 := make([]int, 26)
	arr2 := make([]int, 26)

	x := len(s1)
	y := len(s2)

	for i := 0; i < x; i++ {
		arr1[s1[i]-'a']++
		arr2[s2[i]-'a']++
	}

	for i := x; i < y; i++ {
		if isEqual(arr1, arr2) {
			return true
		}
		arr2[s2[i-x]-'a']--
		arr2[s2[i]-'a']++
	}

	return isEqual(arr1, arr2)
}

func isEqual(arr1, arr2 []int) bool {
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
