package main

import "fmt"

// Create a new type of 'deck'
// which is slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clue"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i ,card)
	}
}
