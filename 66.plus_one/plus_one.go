package plus_one

func PlusOne(digits []int) []int {

	// n := len(digits)

	// for i := n - 1; i >= 0; i-- {
	// 	if digits[i] < 9 {
	// 		digits[i]++
	// 		return digits
	// 	}
	// 	digits[i] = 0
	// }

	// newNumber := make([]int, n+1)
	// newNumber[0] = 1

	// return newNumber

	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			digits[i]++
		}
		if digits[i] == 10 {
			digits[i] = 0
			if i != 0 {
				digits[i-1]++
			} else {
				// Insert 0 at the beginning and set the first digit to 1
				digits = append([]int{0}, digits...)
				digits[0] = 1
			}
		}
	}
	return digits

}
