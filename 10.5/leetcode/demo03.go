package leetcode

import "math"

func maxProfit(prices []int) int {

	minPrice :=prices[0]
	maxProfit:=0
	for i,price := range prices {
		minPrice= int(math.Min(float64(minPrice), float64(price)))
		maxProfit=int(math.Max(float64(maxProfit), float64(prices[i]-minPrice)))
	}
	return maxProfit
}
//贪心