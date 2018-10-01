package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	fmt.Println("Old Card : ", card)

	newCard := getNewCard()
	fmt.Println("New Card : ", newCard)

	fmt.Println("Card Number : ", getCardNumber())

	cardDeck := deck{"Ace of Diamonds", getNewCard()}
	fmt.Println(cardDeck)

	cardDeck = append(cardDeck, "Jack of Diamons")
	fmt.Println(cardDeck)

	fmt.Println("=======Iterating Card=======")
	for a, singleCard := range cardDeck {
		fmt.Println(a, singleCard)
	}

	fmt.Println("=======Iterating Card Through Other File and Function=======")
	deck.print(cardDeck)

	fmt.Println("=======Creating New Deck=======")
	cards := newDeck()
	cards.print()
	fmt.Println("Size of New Deck : ", cards.getSize())

	fmt.Println("=======Make Deal=======")
	var startIndex int
	var endIndex int
	startIndex = rand.Intn(cards.getSize() - 1)
	endIndex = rand.Intn(cards.getSize())
	fmt.Println("Start Index : ", startIndex, " End Index : ", endIndex)
	dealCards := cards.makeDeal(startIndex, endIndex, "first")
	deck.print(dealCards)

	fmt.Println("=======Creating New Deck=======")
	cardsNew := newDeck()
	fmt.Println("=======Making Multiple Deal=======")
	firstDeal, secondDeal := makeMultipleDeal(cardsNew, 4)
	firstDeal.print()
	secondDeal.print()

	fmt.Println("First Deal To String : ", firstDeal.toString())
	fmt.Println("First Deal To Byte Slice : ", toByteSlice(firstDeal.toString()))

	fmt.Println("=======Writing to File=======")
	firstDeal.saveToFile("firstDeal.txt", toByteSlice(firstDeal.toString()))
	secondDeal.saveToFile("secondDeal.txt", toByteSlice(secondDeal.toString()))

	fmt.Println("=======Reading First Deal From File=======")
	firstDeal = convertDataToDeck("firstDeal.txt")
	firstDeal.print()
	fmt.Println("=======Reading Second Deal From File=======")
	secondDeal = convertDataToDeck("secondDeal.txt")
	secondDeal.print()

	fmt.Println("=======Shuffle First Deal=======")
	firstDeal.shuffle().print()
	fmt.Println("=======Shuffle Second Deal=======")
	secondDeal.shuffle().print()
}

func getNewCard() string {
	return "Five Of Diamonds"
}

func getCardNumber() int {
	return 1
}
