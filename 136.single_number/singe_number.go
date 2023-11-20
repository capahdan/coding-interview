package single_number

func SingleNumber(nums []int) int {
	// if len(nums) == 1 {
	// 	return nums[0]
	// }

	// singleMap := make(map[int]int)

	// for _, n := range nums {
	// 	singleMap[n] += 1
	// }
	// singleNumber := 0
	// for key, value := range singleMap {
	// 	if value == 1 {
	// 		singleNumber = key
	// 	}
	// }
	// return singleNumber

	// 	// we sort first the data
	// 	sort.Slice(nums, func(i, j int) bool {
	// 		return nums[i] < nums[j]
	// 	})

	// 	for i := 1; i < len(nums)-1; i += 2 {
	// 		if nums[i] != nums[i-1] {
	// 			return nums[i-1]
	// 		}
	// 	}
	// 	return nums[len(nums)-1]

	//  bitwise solution

	ans := 0
	for _, x := range nums {
		ans ^= x
	}
	return ans

}
