package main

func main() {

	var name string
	name = "Basic Variable"
	println("Value of variable name : ", name)

	name = "Variable Changed"
	println("Value of change variable : ", name)

	directAssignVariable := "Direct Assign Variable"
	println("Direct Assign variable value : ", directAssignVariable)

	var noDataTypeVariable = "No Data Type Variable"
	println("No Data Type Variable : ", noDataTypeVariable)

	var noDataTypeVariable2 = 30
	println("No Data Type Variable 2 : ", noDataTypeVariable2)

	var (
		firstName = "John"
		lastName = "Doe"
	)

	println("First Name : ", firstName, " - Last name : ", lastName)
}
