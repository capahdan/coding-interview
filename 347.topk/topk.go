package topk

// type top struct {
// 	key   int
// 	value int
// }

func TopKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)
	for i := range nums {
		numMap[nums[i]] += 1
	}

	tops := make([]int, len(nums)+1)

	for key, value := range numMap {
		tops[value] = key
	}

	result := []int{}

	for i := len(tops) - 1; i >= 0; i-- {
		if tops[i] != 0 {
			result = append(result, tops[i])
		}

		if len(result) == k {
			return result
		}
	}

	return result
}

//

// 1, 1, 1, 2, 2, 3

// 1-> 3
// 2 -> 2
// 3 -> 1

// make array len k
// array [1,2,3]
// array[:k]
//
