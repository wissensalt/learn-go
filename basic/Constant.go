package main

func main() {
	const FirstName string = "John"
	const LastName string = "Doe"
	const Value uint16 = 1000

	println("First Name : ", FirstName)
	println("Last Name : ", LastName)
	println("Value : ", Value)

	const (
		Sunday = "Sunday"
		Monday = "Monday"
	)

	println(Sunday)
	println(Monday)
}
