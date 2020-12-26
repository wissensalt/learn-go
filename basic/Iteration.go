package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println("loop")
		break
	}

	for a:= 0; a<10; a++ {
		if a%2 == 0 {
			println("Genap")
		} else {
			println("Ganjil")
		}
	}
}
