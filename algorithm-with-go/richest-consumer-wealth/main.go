package main

import "fmt"

// https://leetcode.com/explore/featured/card/the-leetcode-beginners-guide/692/challenge-problems/4423/
func main() {
	customerBankMatrix := [][]int{{7, 1, 3}, {2, 8, 7}, {1, 9, 5}, {10, 1, 0}}
	fmt.Println(GetRichestCostumer(customerBankMatrix))
}

func GetRichestCostumer(consumerBankMatrix [][]int) int {
	result := 0
	var aggregateWealth = make([]int, len(consumerBankMatrix))
	if len(consumerBankMatrix) < 1 {
		return 0
	}

	for i, bank := range consumerBankMatrix {
		for _, j := range bank {
			aggregateWealth[i] += j
		}
	}

	for _, wealth := range aggregateWealth {
		if wealth > result {
			result = wealth
		}
	}

	return result
}
