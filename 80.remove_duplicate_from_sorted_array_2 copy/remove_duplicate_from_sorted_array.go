package remove_duplicate_from_sorted_array_2

// func RemoveDuplicates(nums []int) int {
// 	numberMap := make(map[int]int)
// 	index := 0

// 	for i := range nums {
// 		if val, ok := numberMap[nums[i]]; !ok || val < 2 {
// 			numberMap[nums[i]] += 1
// 			nums[index] = nums[i]
// 			index++
// 		}
// 	}
// 	return index
// }

func RemoveDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	writePtr := 2 // Points to where the next unique element can be written
	for readPtr := 2; readPtr < len(nums); readPtr++ {
		if nums[readPtr] != nums[writePtr-2] {
			nums[writePtr] = nums[readPtr]
			writePtr++
		}
	}

	return writePtr
}

// [0,0,1,1,1,1,2,3,3]

// You can use this Go function to remove duplicates from a sorted integer array while keeping each unique element at most twice.

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
