package jump_game

func JumpGame(nums []int) int {
	goal := len(nums) - 1
	lastIdx := 0
	coverage := 0
	jump := 0

	for i := range nums {
		coverage = max(coverage, i+nums[i])

		if i == lastIdx {
			lastIdx = coverage
			jump++

			if coverage >= goal {
				return jump
			}
		}
	}

	return jump
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// [2,3,1,1,4]
//
//

func jump(nums []int) int {
	totalJumps := 0
	destination := len(nums) - 1
	coverage, lastJumpIdx := 0, 0

	if len(nums) == 1 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		coverage = max(coverage, i+nums[i])

		if i == lastJumpIdx {
			lastJumpIdx = coverage
			totalJumps++

			if coverage >= destination {
				return totalJumps
			}
		}
	}

	return totalJumps
}
