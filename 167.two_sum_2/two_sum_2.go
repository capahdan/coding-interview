package two_sum_2

func TwoSum2(nums []int, target int) []int {
	l, r := 0, len(nums)-1

	for l < r {
		temp := nums[l] + nums[r]
		if temp == target {
			return []int{l + 1, r + 1}
		}
		if temp < target {
			l++
		} else {
			r--
		}
	}

	return []int{}
}
