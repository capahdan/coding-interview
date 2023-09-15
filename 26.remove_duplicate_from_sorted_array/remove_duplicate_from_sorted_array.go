package remove_duplicate_from_sorted_array

func RemoveDuplicates(nums []int) int {
	numberMap := make(map[int]bool)

	var index int

	for i := range nums {
		if ok := numberMap[nums[i]]; !ok {
			numberMap[nums[i]] = true
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
