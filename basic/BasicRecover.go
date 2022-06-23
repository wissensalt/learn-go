package main

import "fmt"

func main() {
	startApp(false)
	fmt.Println("Error Recovered")
}

func finishApp()  {
	fmt.Println("End of application")
	message := recover()
	if message != nil {
		fmt.Println("Message : ", message)
	}
}

func startApp(error bool)  {
	defer finishApp()
	if error {
		panic("Application Error !!")
	}
}