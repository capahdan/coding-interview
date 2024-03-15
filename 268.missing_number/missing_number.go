package missingnumber

import "sort"

func MissingNumber(nums []int) int {

	sort.Ints(nums)

	val := 0
	length := len(nums)

	for val < length {

		if val != nums[val] {
			return val
		}
		if length != nums[length-1] {
			return length
		}

		val++
		length--
	}

	return len(nums)
}
