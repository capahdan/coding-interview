package median

import "sort"

func Median(nums1, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	// if is even so take the tambahkan length++

	// kalo odd ambil di tengah dan ambil nilai 1 lagi di bagi 2

	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	median := length / 2
	if length%2 == 1 {
		// median++

		return float64(nums1[median])
	}
	response := float64(nums1[median]+nums1[median-1]) / float64(2.00)

	return response
}

func Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr)

	left := Sort(arr[:mid])
	right := Sort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	// Sisipkan sisa elemen dari left atau right (jika ada)
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}
