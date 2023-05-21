package ganjil

func Ganjil(nums []int) int {
	var total int
	for i := range nums {
		if nums[i]%2 != 0 {
			total++
		}
	}
	return total
}
