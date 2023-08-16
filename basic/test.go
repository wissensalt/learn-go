package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	for _, value := range mySlice {
		fmt.Println(value)
		if value%4 == 0 {
			break
		}
	}

	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	for key, value := range myMap {
		fmt.Println(key, value)
		if key == "two" {
			break
		}
	}

	newSlice := make([]int, 2, 3)
	newSlice[0] = 1
	newSlice[1] = 2
	fmt.Printf("Len %d , Cap %d\n", len(newSlice), cap(newSlice))

	a, b, c := getCompleteName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	cetakNama := printString
	fmt.Println(cetakNama("John", "Doe"))

	jumlahTotal := sumValues
	fmt.Println("Jumlah Total : ", jumlahTotal(1, 2, 3, 4, 5, 6, 7, 8, 9))

	hitungGenap := calculateEvenNumbers
	fmt.Println("Menghitung hanya angka genap : ", hitungGenap(filterEvenNumber, 1, 2, 3, 4, 5, 6))
}

func filterEvenNumber(number int) int {
	if number%2 == 0 {
		return number
	}

	return 0
}

func calculateEvenNumbers(filter func(input int) int, numbers ...int) int {
	var totalNumber int
	for _, i := range numbers {
		totalNumber += filter(i)
	}

	return totalNumber
}

func printString(one string, two string) (string, string) {
	return one, two
}

func sumValues(values ...int) int {
	var totalValue int
	for _, i := range values {
		totalValue += i
	}

	return totalValue
}

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "John"
	middleName = "Doe"
	lastName = "Does"

	return
}

func learnMap() {
	daysInWeek := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}
	fmt.Println(daysInWeek)
	fmt.Println("First Day : ", daysInWeek[1])
	fmt.Println("Last Day : ", daysInWeek[7])
	// update last day
	daysInWeek[7] = "Off day"
	fmt.Println(daysInWeek)
	delete(daysInWeek, 7)
	fmt.Println(daysInWeek)
	fmt.Println(len(daysInWeek))

	// another way
	employee := make(map[string]string)
	employee["id"] = "123456"
	employee["name"] = "John"
	employee["department"] = "Finance"
	fmt.Println(employee)
}

func learnSlice() {
	var myArray = [...]string{"A", "B", "C", "D", "E"}
	var mySlice = myArray[1:4]
	fmt.Println(mySlice)      //BCD
	fmt.Println(len(mySlice)) //3
	fmt.Println(cap(mySlice)) //4
	var additionalSlice = append(mySlice, "F", "G", "H")
	fmt.Println(additionalSlice)
	fmt.Println(cap(additionalSlice))
	additionalSlice[len(additionalSlice)-1] = "X"
	fmt.Println(additionalSlice)
	fmt.Println(myArray)

	newSlice := make([]string, len(mySlice), cap(mySlice))
	copy(newSlice, mySlice)
	fmt.Println(newSlice)
}
