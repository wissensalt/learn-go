package main

import "fmt"

func main() {
	displayAll(10, 1)
	displayOnlyEven(10, 1)
	displayOnlyOdd(10, 1)
}

func displayAll(n int, counter int) {
	if counter < n {
		if counter%2 == 0 {
			fmt.Println(counter, "is even")
		} else {
			fmt.Println(counter, "is odd")
		}
		counter++
		displayAll(n, counter)
	}
}

func displayOnlyEven(n int, counter int) {
	if counter < n {
		if counter%2 == 0 {
			fmt.Println(counter, "is even")
		}
		counter++
		displayOnlyEven(n, counter)
	}
}

func displayOnlyOdd(n int, counter int) {
	if counter < n {
		if counter%2 != 0 {
			fmt.Println(counter, "is odd")
		}
		counter++
		displayOnlyOdd(n, counter)
	}
}
