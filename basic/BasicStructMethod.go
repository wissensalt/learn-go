package main

import "fmt"

type Employee struct {
	Name, Address string
	Age           int8
}

func (employee Employee) sayHello() {
	fmt.Println("Hello my name is : ", employee.Name)
}

func main() {
	donny := Employee{"Donny", "Jkt", 22}
	donny.sayHello()
}
