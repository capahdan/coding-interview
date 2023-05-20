package palindrome_number

func IsPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 {
		return false
	}
	temp := x
	var reversed int
	for x != 0 {
		remainder := x % 10
		reversed = reversed*10 + remainder
		x = x / 10
	}
	return temp == reversed
}
