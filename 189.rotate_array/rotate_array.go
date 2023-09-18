package rotate_array

// {
// 	name: "test1",
// 	input: Number{
// 		nums: []int{1, 2, 3, 4, 5, 6, 7},
// 		k:    3,
// 	},
// 	expected: []int{5, 6, 7, 1, 2, 3, 4},
// },

func RotateArray(nums []int, k int) []int {
	n := len(nums)
	k %= n

	// reverse the entire row
	reverse(nums, 0, n-1)

	// reverse firt k element

	reverse(nums, 0, k-1)
	// reverse  the remaining arr

	reverse(nums, k, n-1)

	return nums
}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
