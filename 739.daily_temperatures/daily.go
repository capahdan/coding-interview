package daily

func DailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []int{}

	for i, temp := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = idx - i
		}
		stack = append(stack, i)
	}

	return result
}
