package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		// %v - passed variable as argument (len(d))
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

}
