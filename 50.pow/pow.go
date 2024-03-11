package pow

// Create a function pow that takes two argument, a and b. And then it will calculate the a to the power of b. Please do not use built in function to calculate the power

func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}

	if n%2 == 0 {
		return MyPow(x*x, n/2)
	}

	return x * MyPow(x*x, n/2)
}
