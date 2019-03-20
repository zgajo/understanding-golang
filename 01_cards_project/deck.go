package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	// cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	var cardSuits = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cs := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+cs)
		}
	}

	return cards
}

func (d deck) print() {
	// range = iteare over every elemtent in cards
	for i, card := range d {
		fmt.Println(i, card)
	}
}
