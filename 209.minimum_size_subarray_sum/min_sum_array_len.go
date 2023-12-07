package min_sum_array_len

import (
	"math"
)

func MinSubArrayLen(target int, nums []int) int {
	minLen := math.MaxInt
	var sum, start int

	for end := 0; end < len(nums); end++ {
		sum += nums[end]

		for sum >= target {
			if end-start+1 < minLen {
				minLen = end - start + 1
			}
			sum -= nums[start]
			start++
		}
	}

	if minLen == math.MaxInt {
		return 0
	}

	return minLen
}

// 2, 3, 1, 2, 4, 3

//  sum = 2+3+1+2

// sum -= 2
// sum = 3 +1 +2 +4

// sum -= 3
// sum = 1 + 2 + 4

// sum -= 1
// sum = 2+4+3

// sum -= 2
// sum = 4+3
