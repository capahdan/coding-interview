package pow

// Create a function pow that takes two argument, a and b. And then it will calculate the a to the power of b. Please do not use built in function to calculate the power

func Power(a float64, b int) float64 {

	if b == 0 {
		return 1
	}

	return a * Power(a, b-1)
}
