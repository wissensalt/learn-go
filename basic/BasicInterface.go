package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func sayTheName(name HasName) {
	fmt.Println("Hello : ", name.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

type Car struct {
	Name string
}

func (car Car) GetName() string {
	return car.Name
}

func main() {
	john := Person{"John"}
	sayTheName(john)

	lamborghini := Car{"Lamborghini"}
	sayTheName(lamborghini)
	fmt.Println(lamborghini.GetName())
}
