package main

import "fmt"

func main() {
	seconds := 20000
	fmt.Println("Input Seconds", seconds)
	hour := seconds / 3600
	hourMod := seconds % 3600
	min := hourMod / 60
	seconds = hourMod % 60

	fmt.Println("Hours ", hour)
	fmt.Println("Mins ", min)
	fmt.Println("seconds ", seconds)
}
