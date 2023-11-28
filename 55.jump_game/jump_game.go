package jump_game

import "math"

func CanJump(nums []int) bool {
	dis := 0
	for i := 0; i <= dis; i++ {
		dis = int(math.Max(float64(dis), float64(i+nums[i])))
		if dis >= len(nums)-1 {
			return true
		}
	}
	return false
}
