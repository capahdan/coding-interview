package merge_sorted_array

// func Merge(nums1 []int, m int, nums2 []int, n int) []int {
// 	numsResult := nums1[:m]

// 	numsResult = append(numsResult, nums2...)

// 	sort.Slice(numsResult, func(i, j int) bool {
// 		return numsResult[i] < numsResult[j]
// 	})
// 	return numsResult
// }

//

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	for p2 >= 0 && p1 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}

	return nums1
}
