package main

import (
	"fmt"

	"./base"
	"./object"
)

func main() {
	var baseObject base.Calculate

	baseObject = object.Square{2}
	fmt.Println("===== Square")
	fmt.Println("ResultA      :", baseObject.GetResultA())
	fmt.Println("ResultB      :", baseObject.GetResultB())

	baseObject = object.Circle{2}
	fmt.Println("===== CIRCLE")
	fmt.Println("JARIJARI     :", baseObject.(object.Circle).GetJariJari())
	fmt.Println("ResultA      :", baseObject.GetResultA())
	fmt.Println("ResultB      :", baseObject.GetResultB())
}
