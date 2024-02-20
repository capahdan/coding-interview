package way_to_sum

import "fmt"

// Solution finds the first missing positive integer in an array of integers.
func WayToSum(total, k int) int {
	// Initialize a slice to store the counts
	dp := make([]int, total+1)
	dp[0] = 1 // There's one way to achieve a total weight of 0

	// Iterate through each weight from 1 to k
	for weight := 1; weight <= k; weight++ {
		// Update the count for each total weight achievable with the current weight
		for j := weight; j <= total; j++ {
			dp[j] += dp[j-weight]

			fmt.Println("=====", dp)

			fmt.Println("j", j, "weight", weight)
		}
	}

	return dp[total]
}

//1 1 1 1 1
//
