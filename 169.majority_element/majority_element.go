package majority_element

func MajorityElement(nums []int) int {
	var majority int

	half := len(nums) / 2

	if len(nums) <= 1 {
		return nums[0]
	}

	numberMap := make(map[int]int)

	for i := range nums {
		if val, ok := numberMap[nums[i]]; ok && val >= half {
			majority = nums[i]
		} else {
			numberMap[nums[i]] += 1
		}

	}

	return majority
}
