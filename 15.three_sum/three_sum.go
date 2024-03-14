package three_sum

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {

	tempMap := map[string][]int{}
	result := [][]int{}
	// temp := make([]int, len(nums)-1)

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		target := 0 - nums[i]

		l := i + 1
		r := len(nums) - 1
		for l < r {

			temp := nums[l] + nums[r]
			if target == temp {
				if i != l && i != r && l != r {
					tempMap[fmt.Sprintf("%d-%d-%d", nums[i], nums[l], nums[r])] = []int{nums[i], nums[l], nums[r]}
				}
			}
			if temp < target {
				l++
			} else {
				r--
			}
		}
	}

	for _, val := range tempMap {
		result = append(result, val)
	}

	return result
}
