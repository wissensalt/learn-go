package main

import "fmt"

var result bool

func main() {
	var val = 10
	var check = checkPrime(2, val)

	if !check {
		fmt.Println(val, "is a prime")
	} else {
		fmt.Println(val, "is NOT a prime")
	}
}

func checkPrime(inc int, val int) bool {
	if val == 2 {
		result = false
	}
	if inc < val {
		if val%inc == 0 {
			result = true
		} else {
			inc++
			checkPrime(inc, val)
			return result
		}
	} else {
		result = false
	}
	return result
}
