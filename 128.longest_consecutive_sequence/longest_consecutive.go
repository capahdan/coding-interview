package longest_consecutive

// my solution
// func LongestConsecutive(nums []int) int {
// 	var result int

// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	if len(nums) == 1 {
// 		return 1
// 	}
// 	sort.Ints(nums)

// 	result++
// 	// expected value yang kita inginkan
// 	seqMap := make(map[int]int)

// 	seqMap[nums[0]+1] = nums[0]

// 	for i := 1; i < len(nums); i++ {
// 		if _, ok := seqMap[nums[i]]; ok {
// 			result++
// 		}
// 		seqMap[nums[i]+1] = seqMap[nums[i]]
// 	}

// 	return result
// }

// func LongestConsecutive(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	maxLength := 0
// 	numMap := make(map[int]int)

// 	for _, num := range nums {
// 		if _, exists := numMap[num]; !exists {
// 			left := numMap[num-1]
// 			right := numMap[num+1]

// 			// sum is the length of the sequence num is in
// 			sum := left + right + 1
// 			numMap[num] = sum

// 			// update the max length
// 			if sum > maxLength {
// 				maxLength = sum
// 			}

// 			// extend the length to the boundaries of the sequence
// 			numMap[num-left] = sum
// 			numMap[num+right] = sum
// 		}
// 	}

// 	return maxLength
// }

// urutkan nums secara asc
//
// e.g = 100,4,200, 1, 3, 2

// sorted = 1,2,3,4,100,200

//
// seqMap 1->2
//        2->3
//

//

func LongestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	best := 0
	for num := range numSet {
		if !numSet[num-1] {
			y := num + 1
			for numSet[y] {
				y++
			}
			if y-num > best {
				best = y - num
			}
		}
	}
	return best
}

// ubah semua number jadi set
// lalu iterasi setiap set
// jika nilai set  value-1 tidak ada di dalam set
//     iterasi set  value+1
//     jika ada kita update
//     nilai bestnya
