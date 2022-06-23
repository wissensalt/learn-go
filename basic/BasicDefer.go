package main

import "fmt"

func main() {
	runApplication(0)
}

func runApplication(value int)  {
	defer logging()
	result := 1 / value
	fmt.Println("Result : ", result)
}

func logging() {
	fmt.Println("Defer Logging at the end of statement, It must be shown even though there is an error !!")
}