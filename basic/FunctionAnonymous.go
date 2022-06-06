package main

import "fmt"

type Blacklist func(val string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("John", blacklist)
	registerUser("Doe", func(val string) bool {
		return val == "root"
	})
}