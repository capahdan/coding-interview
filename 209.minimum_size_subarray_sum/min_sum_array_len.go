package min_sum_array_len

import "math"

func MinSubArrayLen(target int, nums []int) int {

	minLength := math.MaxInt32
	sum, start := 0, 0

	for end := 0; end < len(nums); end++ {
		sum += nums[end]

		for sum >= target {
			if end-start+1 < minLength {
				minLength = end - start + 1
			}
			sum -= nums[start]
			start++
		}
	}

	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}

// min := Math.Min()

// 2,3,1,2,4,3

// min := []int

//
