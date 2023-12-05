package min_sum_array_len

import "math"

func MinSubArrayLen(target int, nums []int) int {

	minLen := math.MaxInt

	for i := 0; i < len(nums); i++ {
		totalNum := 0
		times := 0
		for j := 0; j < len(nums); j++ {
			totalNum += nums[j]
			times++

			if totalNum >= target {
				minLen = int(math.Min(float64(minLen), float64(times)))
				totalNum = 0
				times = 0
			}
		}
	}

	if minLen == math.MaxInt {
		return 0
	}

	return minLen
}

// min := Math.Min()

// 2,3,1,2,4,3

// min := []int

//
