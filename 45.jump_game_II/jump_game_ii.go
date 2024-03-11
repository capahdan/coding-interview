package jump_game

import "math"

func JumpGame(nums []int) int {

	// x := 0
	n := len(nums) - 1
	mov := math.MaxInt32

	for i := range nums {
		for j := nums[i]; i >= 1; j-- {
			jump := 0
			k := nums[j]
			for k <= n {
				jump++
				k += nums[k]
			}

			mov = min(mov, jump)
		}

	}

	return mov
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// [2, 3, 1, 1, 4]
//  i
//  j

// 0
// 2
