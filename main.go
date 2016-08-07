package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	d := NewStdDeck()
	allHands := GenAllHands(&d)

	runTime := time.Since(startTime)

	fmt.Println("New deck: ", d.String())
	fmt.Println("Number of hands: ", len(allHands))
	fmt.Println("Ran in ", runTime)
}
