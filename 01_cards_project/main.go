package main

import "fmt"

func main() {
	// Two ways to define variables,
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // only to define new variable

	// array vs slice
	// Array - fixed length list of things
	// Slice - array that cang grow and shrink
	// slice
	cards := []string{"Ace", newCard()}

	// range = iteare over every elemtent in cards
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Test"
}
