package contain_duplicate

import "sort"

func ContainsDuplicate(nums []int) bool {

	sort.Ints(nums)
	mapNum := make(map[int]bool)

	for i := range nums {
		if _, ok := mapNum[nums[i]]; ok {
			return true
		}
		mapNum[nums[i]] = true
	}

	return false
}
