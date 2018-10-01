package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Curly"}
	cardValues := []string{"Ace", "Two", "Three"}
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" Of "+suite)
		}
	}
	return cards
}

func (d deck) print() {
	for a, card := range d {
		fmt.Println(a, card)
	}
}

func (d deck) getSize() int {
	return len(d)
}

func (d deck) makeDeal(startIndex int, endIndex int, action string) []string {
	if action == "first" {
		fmt.Println("Performing First Slice ", startIndex)
		return d[startIndex:]
	} else if action == "last" {
		fmt.Println("Performing Last Slice ", endIndex)
		return d[:endIndex]
	} else {
		fmt.Println("Performing Both Slice ")
		return d[startIndex:endIndex]
	}
}

func makeMultipleDeal(d deck, index int) (deck, deck) {
	return d[index:], d[:index]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func toByteSlice(data string) []byte {
	return []byte(data)
}

func (d deck) saveToFile(pFileName string, pRawData []byte) error {
	return ioutil.WriteFile(pFileName, pRawData, 0666)
}

func convertDataToDeck(pFileName string) deck {
	rawData, err := ioutil.ReadFile(pFileName)
	if err != nil {
		fmt.Println("Error Convert Data To Deck ", err)
		os.Exit(1)
		return nil
	} else {
		return strings.Split(string(rawData), ",")
	}
}

func (d deck) shuffle() deck {
	// rand.Intn(len(d) - 1)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for a := range d {
		newPosition := r.Intn(len(d) - 1)
		d[a], d[newPosition] = d[newPosition], d[a]
	}
	return d
}
