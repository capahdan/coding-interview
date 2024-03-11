package container

// input:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
// output:  49,

// method 2 pointer yang mana pointer pertama

// maxValue

// [1, 8, 6, 2, 5, 4, 8, 3, 7]
//  a                       b

// 1 -> 0
// 8 -> 1
// ....
//  7 -> 8

func MaxArea(height []int) int {
	maxValue := 0
	a := 0
	b := len(height) - 1

	for a != b {
		val := less(height[a], height[b])

		if val*(b-a) > maxValue {
			maxValue = val * (b - a)
		}
		if height[b] > height[a] {
			a++
		} else {
			b--
		}

	}

	return maxValue
}

func less(a, b int) int {
	if a < b {
		return a
	}
	return b
}
