package remove_duplicate_from_sorted_array_2

func RemoveDuplicates(nums []int) int {
	numberMap := make(map[int]int)
	index := 0

	for i := range nums {
		if val, ok := numberMap[nums[i]]; !ok || val < 2 {
			numberMap[nums[i]] += 1
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

// [1,1,1,2,2,3]
//expected 5

// first iteration
// i = 0
// nums =1
//  numberMap[1] = 1

// second iteration
// i =1
//  nums =1
//  numMap[1]=2

// third iteration

// i =2
//  nums =1
// numMap
