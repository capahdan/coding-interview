package three_sum

import "fmt"

func ThreeSum(nums []int) [][]int {
	mapArr := make(map[string][]int)

	for i := len(nums) - 1; i >= 2; i-- {

		if nums[i]+nums[i-1]+nums[i-2] == 0 {
			mapArr[fmt.Sprintf("%d-%d-%d", i, i-1, i-2)] = []int{nums[i], nums[i-1], nums[i-2]}
		}

	}

	result := [][]int{}
	for _, val := range mapArr {
		result = append(result, val)
	}

	return result
}
