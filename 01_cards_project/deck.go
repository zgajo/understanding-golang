package main

import "fmt"

type deck []string

func (d deck) print() {
	// range = iteare over every elemtent in cards
	for i, card := range d {
		fmt.Println(i, card)
	}
}
