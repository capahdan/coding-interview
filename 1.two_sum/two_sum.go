package two_sum

func TwoSum(nums []int, target int) []int {
	numberMap := make(map[int]int)

	for i := range nums {
		dif := target - nums[i]
		if val, ok := numberMap[dif]; ok {
			return []int{val, i}
		}
		numberMap[nums[i]] = i
	}
	return nil
}
