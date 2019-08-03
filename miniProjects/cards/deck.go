package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	color := []string{"heart", "diamond"}
	num := []string{"ace", "two", "three"}

	// var cards deck = {"cunt"}
	for _, c := range color {
		for _, n := range num {
			cards = append(cards, n+" of "+c)
		}
	}
	return cards
}

func (d deck) printCards() {
	for _, d := range d {
		fmt.Println(d)
	}
}
