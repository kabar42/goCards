package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	c := Card{Ace, Spades}
	fmt.Println("New card: ", c.Rank, c.Suit)

	d := NewStdDeck()
	fmt.Println("New deck: ", d.Cards)
}
