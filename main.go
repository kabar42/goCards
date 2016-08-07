package main

import "fmt"

func main() {
	d := NewStdDeck()
	fmt.Println("New deck: ", d.String())

	h := NewHand()
	allHands := make([]Hand, 0)
	GenAllHands(&d, &h, &allHands)

	fmt.Println("Number of hands: ", len(allHands))
}
