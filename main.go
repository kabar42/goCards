package main

import "fmt"

func main() {
	d := NewStdDeck()
	fmt.Println("New deck: ", d.String())

	allHands := GenAllHands(&d)

	fmt.Println("Number of hands: ", len(allHands))
}
