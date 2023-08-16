package main

import "fmt"

// https://leetcode.com/explore/featured/card/the-leetcode-beginners-guide/692/challenge-problems/4422/
func main() {
	nums := []int{3, 1, 2, 10, 1}
	var result = SumOneDArray(nums)
	fmt.Println(result)
}

func SumOneDArray(nums []int) []int {
	var result = make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			result[i] = num
		} else {
			result[i] = result[i-1] + num
		}
	}

	return result
}
