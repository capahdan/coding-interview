package best_time_to_buy_and_sell_stock

import "math"

func MaxProfit(prices []int) int {

	// {
	// name: "test1",
	// input: Number{
	// 	prices: []int{7, 1, 5, 3, 6, 4},
	// },
	// expected: 5,

	// index := 0
	// lowestNumber := prices[index]
	// idxOfLowest := 0
	// arrayLength := len(prices) - 1

	// maxProfit := 0

	// for i := 0; i < len(prices); i++ {
	// 	if prices[i] < lowestNumber && i != arrayLength && prices[i] != 0 {
	// 		lowestNumber = prices[i]
	// 		idxOfLowest = i
	// 	}
	// }

	// for j := idxOfLowest; j < len(prices); j++ {
	// 	dif := prices[j] - lowestNumber

	// 	if dif > maxProfit {
	// 		maxProfit = dif
	// 	}
	// }

	// return maxProfit

	// chat gpt solution

	lsf := math.MaxInt
	op := 0
	pist := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < lsf {
			lsf = prices[i]
		}
		pist = prices[i] - lsf
		if op < pist {
			op = pist
		}
	}
	return op

}
