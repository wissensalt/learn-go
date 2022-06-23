package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var john Customer
	john.Name = "John"
	john.Address = "Jkt"
	john.Age = 25

	fmt.Println(john)

	foo := Customer{
		Name:    "Foo",
		Address: "Bdg",
		Age:     24,
	}

	fmt.Println(foo)

	bar := Customer{"Bar", "Sby", 26}

	fmt.Println(bar)
}
