package main

import "fmt"

func main() {
	for a := 2; a <= 10; a++ {
		check := false
		for b := 2; b < a; b++ {
			if a%b == 0 {
				check = true
			}
		}
		if !check {
			fmt.Println(a, "is prime value")
		}
	}
}
