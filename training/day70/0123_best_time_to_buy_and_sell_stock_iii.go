package day70

import "sort"

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	profit := make([]int, 0)
	for i := 1; i < len(prices); i++ {
		profit = append(profit, prices[i]-prices[i-1])
	}

	result := make([]int, 0)
	index := -1
	for i := 0; i < len(profit); i++ {
		if profit[i] < 0 {
			result = append(result, 0)
			index++
		} else if i == 0 {
			result = append(result, profit[i])
			index++
		} else {
			result[index] += profit[i]
		}
	}
	sort.Ints(result)

	if len(result) == 1 {
		return result[0]
	}
	return result[len(result)-1] + result[len(result)-2]
}
