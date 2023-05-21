package two_sum

func TwoSum(nums []int, target int) []int {
	data := make(map[int]int)
	for i := range nums {
		searced := target - nums[i]
		if index, exist := data[searced]; exist {
			return []int{index, i}
		}
		data[nums[i]] = i
	}
	return nil
}
