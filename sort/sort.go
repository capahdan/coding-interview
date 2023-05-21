package sort

func Sort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] < nums[j+1] {
				swap(&nums[j], &nums[j+1])
			}
		}
	}
	return nums
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
