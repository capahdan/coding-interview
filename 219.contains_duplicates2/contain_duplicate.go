package contain_duplicate_2

import "fmt"

func ContainsNearbyDuplicate(nums []int, k int) bool {
	// map [int]int ,

	numMap := make(map[int]int)

	for index := range nums {
		v, ok := numMap[nums[index]]
		fmt.Println(v, ok)
		if v, ok := numMap[nums[index]]; ok && index-v <= k {
			return true
		}
		numMap[nums[index]] = index
	}

	return false
}
