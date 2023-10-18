package summary_range

import (
	"strconv"
)

type Stack struct {
	num []int
}

func SummaryRanges(nums []int) []string {
	result := []string{}

	for i := 0; i < len(nums); {
		start := nums[i]
		for i+1 < len(nums) && nums[i]+1 == nums[i+1] {
			i++
		}

		if start != nums[i] {
			result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i]))
		} else {
			result = append(result, strconv.Itoa(start))
		}
		i++
	}

	return result
}
