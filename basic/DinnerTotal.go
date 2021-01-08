package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Usage : DinnerTotal <Meal Total> <Tip Amount>")
		fmt.Println("Example : ./DinnerTotal 5000 100")
	} else {
		if len(args) != 2 {
			fmt.Println("You must 2 arguments!. Type /help for help")
		} else {
			mealTotal, _ := strconv.ParseFloat(args[0], 32)
			tipAmount, _ := strconv.ParseFloat(args[1], 32)
			fmt.Printf("Your meal total will be %.2f\n",
				calculateTotal(float32(mealTotal), float32(tipAmount)))
		}
	}
}

func calculateTotal(mealTotal float32, tipAmount float32) float32 {
	return mealTotal + (mealTotal * (tipAmount / 100))
}
