package main

import "fmt"

func main() {
	runApp(true)
}

func endApp() {
	fmt.Println("Application Finished")
}

func runApp(error bool)  {
	defer endApp()
	if error {
		panic("Application Error!!")
	}

	fmt.Println("Application Run Properly")
}